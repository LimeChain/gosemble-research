package dev

import (
	"errors"
	"fmt"
	"sync"

	"github.com/radkomih/gosemble/utils"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

var (
	ErrCodeEmpty              = errors.New("code is empty")
	ErrWASMDecompress         = errors.New("wasm decompression failed")
	ErrInstanceIsStopped      = errors.New("instance is stopped")
	ErrExportFunctionNotFound = errors.New("export function not found")
)

type Instance struct {
	vm       wasmer.Instance
	ctx      *Context
	isClosed bool
	// codeHash common.Hash
	mutex sync.Mutex
}

type Context struct {
	// Storage         Storage
	Allocator *FreeingBumpHeapAllocator
	// Keystore  *keystore.GlobalKeystore
	// Validator bool
	// NodeStorage     NodeStorage
	// Network         BasicNetwork
	// Transaction     TransactionState
	// SigVerifier     *crypto.SignatureVerifier
	// OffchainHTTPSet *offchain.HTTPSet
	// Version         Version
}

// Config is the configuration used to create a Wasmer runtime instance.
type Config struct {
	// Storage     runtime.Storage
	// Keystore    *keystore.GlobalKeystore
	// LogLvl log.Level
	// Role        common.Roles
	// NodeStorage runtime.NodeStorage
	// Network     runtime.BasicNetwork
	// Transaction runtime.TransactionState
	// CodeHash common.Hash
	// testVersion *runtime.Version
}

func setupVM(code []byte) (instance wasmer.Instance, allocator *FreeingBumpHeapAllocator, err error) {
	if len(code) == 0 {
		return instance, nil, ErrCodeEmpty
	}

	// code, err = decompressWasm(code)
	// if err != nil {
	// 	// Note the sentinel error is wrapped here since the ztsd Go library
	// 	// does not return any exported sentinel errors.
	// 	return instance, nil, fmt.Errorf("%w: %s", ErrWASMDecompress, err)
	// }

	imports, err := importsNodeRuntime()
	if err != nil {
		return instance, nil, fmt.Errorf("creating node runtime imports: %w", err)
	}

	// Provide importable memory for newer runtimes
	// TODO: determine memory descriptor size that the runtime wants from the wasm.
	// should be doable w/ wasmer 1.0.0. (#1268)
	memory, err := wasmer.NewMemory(23, 0)
	if err != nil {
		return instance, nil, fmt.Errorf("creating web assembly memory: %w", err)
	}

	_, err = imports.AppendMemory("memory", memory)
	if err != nil {
		return instance, nil, fmt.Errorf("appending memory to imports: %w", err)
	}

	// Instantiates the WebAssembly module.
	instance, err = wasmer.NewInstanceWithImports(code, imports)
	if err != nil {
		return instance, nil, fmt.Errorf("creating web assembly instance: %w", err)
	}

	// Assume imported memory is used if runtime does not export any
	if !instance.HasMemory() {
		instance.Memory = memory
	}

	// TODO: get __heap_base exported value from runtime.
	// wasmer 0.3.x does not support this, but wasmer 1.0.0 does (#1268)
	heapBase := DefaultHeapBase

	allocator = NewAllocator(instance.Memory, heapBase)

	return instance, allocator, nil
}

func NewInstance(code []byte, cfg Config) (instance *Instance, err error) {
	// logger.Patch(log.SetLevel(cfg.LogLvl), log.SetCallerFunc(true))

	wasmInstance, allocator, err := setupVM(code)
	if err != nil {
		return nil, fmt.Errorf("setting up VM: %w", err)
	}

	runtimeCtx := &Context{
		// Storage:         cfg.Storage,
		Allocator: allocator,
		// Keystore:        cfg.Keystore,
		// Validator:       cfg.Role == common.AuthorityRole,
		// NodeStorage:     cfg.NodeStorage,
		// Network:         cfg.Network,
		// Transaction:     cfg.Transaction,
		// SigVerifier:     crypto.NewSignatureVerifier(logger),
		// OffchainHTTPSet: offchain.NewHTTPSet(),
	}
	wasmInstance.SetContextData(runtimeCtx)

	instance = &Instance{
		vm:  wasmInstance,
		ctx: runtimeCtx,
		// codeHash: cfg.CodeHash,
	}

	// if cfg.testVersion != nil {
	// 	instance.ctx.Version = *cfg.testVersion
	// } else {
	// 	instance.ctx.Version, err = instance.version()
	// 	if err != nil {
	// 		instance.close()
	// 		return nil, fmt.Errorf("getting instance version: %w", err)
	// 	}
	// }

	wasmInstance.SetContextData(instance.ctx)

	return instance, nil
}

func (in *Instance) Exec(function string, data []byte) (result []byte, err error) {
	in.mutex.Lock()
	defer in.mutex.Unlock()

	if in.isClosed {
		return nil, ErrInstanceIsStopped
	}

	dataLength := uint32(len(data))
	inputPtr, err := in.ctx.Allocator.Allocate(dataLength)
	if err != nil {
		return nil, fmt.Errorf("allocating input memory: %w", err)
	}

	defer in.ctx.Allocator.Clear()

	// Store the data into memory
	memory := in.vm.Memory.Data()
	copy(memory[inputPtr:inputPtr+dataLength], data)

	// TODO: remove it, only for testing
	_startFunc, ok := in.vm.Exports["_start"]
	if !ok {
		return nil, fmt.Errorf("running runtime _start function: %w", ok)
	}
	_startFunc()

	runtimeFunc, ok := in.vm.Exports[function]
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrExportFunctionNotFound, function)
	}

	wasmValue, err := runtimeFunc(int32(inputPtr), int32(dataLength))
	if err != nil {
		return nil, fmt.Errorf("running runtime function: %w", err)
	}

	outputPtr, outputLength := utils.Int64ToOffsetAndSize(wasmValue.ToI64())
	memory = in.vm.Memory.Data() // call Data() again to get larger slice

	fmt.Printf("%s", memory)
	fmt.Printf("\n\nptr-size: %v, bytes: %0x \n", wasmValue, memory[outputPtr:outputPtr+outputLength])
	return memory[outputPtr : outputPtr+outputLength], nil
}
