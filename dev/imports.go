package dev

// #include <stdlib.h>
//
// extern void ext_logging_log_version_1(void *context, int32_t level, int64_t target, int64_t msg);
// extern int32_t ext_logging_max_level_version_1(void *context);
//
// extern void ext_sandbox_instance_teardown_version_1(void *context, int32_t a);
// extern int32_t ext_sandbox_instantiate_version_1(void *context, int32_t a, int64_t b, int64_t c, int32_t d);
// extern int32_t ext_sandbox_invoke_version_1(void *context, int32_t a, int64_t b, int64_t c, int32_t d, int32_t e, int32_t f);
// extern int32_t ext_sandbox_memory_get_version_1(void *context, int32_t a, int32_t b, int32_t c, int32_t d);
// extern int32_t ext_sandbox_memory_new_version_1(void *context, int32_t a, int32_t b);
// extern int32_t ext_sandbox_memory_set_version_1(void *context, int32_t a, int32_t b, int32_t c, int32_t d);
// extern void ext_sandbox_memory_teardown_version_1(void *context, int32_t a);
//
// extern int32_t ext_crypto_ed25519_generate_version_1(void *context, int32_t a, int64_t b);
// extern int64_t ext_crypto_ed25519_public_keys_version_1(void *context, int32_t a);
// extern int64_t ext_crypto_ed25519_sign_version_1(void *context, int32_t a, int32_t b, int64_t c);
// extern int32_t ext_crypto_ed25519_verify_version_1(void *context, int32_t a, int64_t b, int32_t c);
// extern int32_t ext_crypto_finish_batch_verify_version_1(void *context);
// extern int64_t ext_crypto_secp256k1_ecdsa_recover_version_1(void *context, int32_t a, int32_t b);
// extern int64_t ext_crypto_secp256k1_ecdsa_recover_version_2(void *context, int32_t a, int32_t b);
// extern int64_t ext_crypto_secp256k1_ecdsa_recover_compressed_version_1(void *context, int32_t a, int32_t b);
// extern int64_t ext_crypto_secp256k1_ecdsa_recover_compressed_version_2(void *context, int32_t a, int32_t b);
// extern int32_t ext_crypto_ecdsa_verify_version_2(void *context, int32_t a, int64_t b, int32_t c);
// extern int32_t ext_crypto_sr25519_generate_version_1(void *context, int32_t a, int64_t b);
// extern int64_t ext_crypto_sr25519_public_keys_version_1(void *context, int32_t a);
// extern int64_t ext_crypto_sr25519_sign_version_1(void *context, int32_t a, int32_t b, int64_t c);
// extern int32_t ext_crypto_sr25519_verify_version_1(void *context, int32_t a, int64_t b, int32_t c);
// extern int32_t ext_crypto_sr25519_verify_version_2(void *context, int32_t a, int64_t b, int32_t c);
// extern void ext_crypto_start_batch_verify_version_1(void *context);
//
// extern int32_t ext_trie_blake2_256_root_version_1(void *context, int64_t a);
// extern int32_t ext_trie_blake2_256_ordered_root_version_1(void *context, int64_t a);
// extern int32_t ext_trie_blake2_256_ordered_root_version_2(void *context, int64_t a, int32_t b);
// extern int32_t ext_trie_blake2_256_verify_proof_version_1(void *context, int32_t a, int64_t b, int64_t c, int64_t d);
//
// extern int64_t ext_misc_runtime_version_version_1(void *context, int64_t a);
// extern void ext_misc_print_hex_version_1(void *context, int64_t a);
// extern void ext_misc_print_num_version_1(void *context, int64_t a);
// extern void ext_misc_print_utf8_version_1(void *context, int64_t a);
//
// extern void ext_default_child_storage_clear_version_1(void *context, int64_t a, int64_t b);
// extern int64_t ext_default_child_storage_get_version_1(void *context, int64_t a, int64_t b);
// extern int64_t ext_default_child_storage_next_key_version_1(void *context, int64_t a, int64_t b);
// extern int64_t ext_default_child_storage_read_version_1(void *context, int64_t a, int64_t b, int64_t c, int32_t d);
// extern int64_t ext_default_child_storage_root_version_1(void *context, int64_t a);
// extern void ext_default_child_storage_set_version_1(void *context, int64_t a, int64_t b, int64_t c);
// extern void ext_default_child_storage_storage_kill_version_1(void *context, int64_t a);
// extern int32_t ext_default_child_storage_storage_kill_version_2(void *context, int64_t a, int64_t b);
// extern int64_t ext_default_child_storage_storage_kill_version_3(void *context, int64_t a, int64_t b);
// extern void ext_default_child_storage_clear_prefix_version_1(void *context, int64_t a, int64_t b);
// extern int32_t ext_default_child_storage_exists_version_1(void *context, int64_t a, int64_t b);
//
// extern void ext_allocator_free_version_1(void *context, int32_t a);
// extern int32_t ext_allocator_malloc_version_1(void *context, int32_t a);
//
// extern int32_t ext_hashing_blake2_128_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_blake2_256_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_keccak_256_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_sha2_256_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_twox_256_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_twox_128_version_1(void *context, int64_t a);
// extern int32_t ext_hashing_twox_64_version_1(void *context, int64_t a);
//
// extern void ext_offchain_index_set_version_1(void *context, int64_t a, int64_t b);
// extern int32_t ext_offchain_is_validator_version_1(void *context);
// extern void ext_offchain_local_storage_clear_version_1(void *context, int32_t a, int64_t b);
// extern int32_t ext_offchain_local_storage_compare_and_set_version_1(void *context, int32_t a, int64_t b, int64_t c, int64_t d);
// extern int64_t ext_offchain_local_storage_get_version_1(void *context, int32_t a, int64_t b);
// extern void ext_offchain_local_storage_set_version_1(void *context, int32_t a, int64_t b, int64_t c);
// extern int64_t ext_offchain_network_state_version_1(void *context);
// extern int32_t ext_offchain_random_seed_version_1(void *context);
// extern int64_t ext_offchain_submit_transaction_version_1(void *context, int64_t a);
// extern int64_t ext_offchain_timestamp_version_1(void *context);
// extern void ext_offchain_sleep_until_version_1(void *context, int64_t a);
// extern int64_t ext_offchain_http_request_start_version_1(void *context, int64_t a, int64_t b, int64_t c);
// extern int64_t ext_offchain_http_request_add_header_version_1(void *context, int32_t a, int64_t k, int64_t v);
//
// extern void ext_storage_append_version_1(void *context, int64_t a, int64_t b);
// extern int64_t ext_storage_changes_root_version_1(void *context, int64_t a);
// extern void ext_storage_clear_version_1(void *context, int64_t a);
// extern void ext_storage_clear_prefix_version_1(void *context, int64_t a);
// extern int64_t ext_storage_clear_prefix_version_2(void *context, int64_t a, int64_t b);
// extern void ext_storage_commit_transaction_version_1(void *context);
// extern int32_t ext_storage_exists_version_1(void *context, int64_t a);
// extern int64_t ext_storage_get_version_1(void *context, int64_t a);
// extern int64_t ext_storage_next_key_version_1(void *context, int64_t a);
// extern int64_t ext_storage_read_version_1(void *context, int64_t a, int64_t b, int32_t c);
// extern void ext_storage_rollback_transaction_version_1(void *context);
// extern int64_t ext_storage_root_version_1(void *context);
// extern int64_t ext_storage_root_version_2(void *context, int32_t a);
// extern void ext_storage_set_version_1(void *context, int64_t a, int64_t b);
// extern void ext_storage_start_transaction_version_1(void *context);
//
// extern void ext_transaction_index_index_version_1(void *context, int32_t a, int32_t b, int32_t c);
// extern void ext_transaction_index_renew_version_1(void *context, int32_t a, int32_t b);
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/radkomih/gosemble/utils"
	"github.com/wasmerio/go-ext-wasm/wasmer"
)

//export ext_allocator_malloc_version_1
func ext_allocator_malloc_version_1(context unsafe.Pointer, size C.int32_t) C.int32_t {
	// fmt.Printf("executing malloc with size %d...", int64(size))

	instanceContext := wasmer.IntoInstanceContext(context)
	ctx := instanceContext.Data().(*Context)

	// Allocate memory
	res, err := ctx.Allocator.Allocate(uint32(size))
	if err != nil {
		fmt.Errorf("failed to allocate memory: %s", err)
		panic(err)
	}

	return C.int32_t(res)
}

//export ext_allocator_free_version_1
func ext_allocator_free_version_1(context unsafe.Pointer, addr C.int32_t) {
	// fmt.Printf("executing free...")

	instanceContext := wasmer.IntoInstanceContext(context)
	runtimeCtx := instanceContext.Data().(*Context)

	// Deallocate memory
	err := runtimeCtx.Allocator.Deallocate(uint32(addr))

	if err != nil {
		fmt.Errorf("failed to free memory: %s", err)
	}
}

//export ext_storage_set_version_1
func ext_storage_set_version_1(context unsafe.Pointer, keySpan, valueSpan C.int64_t) {
	fmt.Println("executing...")

	instanceContext := wasmer.IntoInstanceContext(context)
	ctx := instanceContext.Data().(*Context)
	storage := ctx.Storage
	key := asMemorySlice(instanceContext, keySpan)
	value := asMemorySlice(instanceContext, valueSpan)

	cp := make([]byte, len(value))
	copy(cp, value)

	fmt.Println("key 0x%x has value 0x%x", key, value)

	storage.Set(key, cp)
}

// Convert 64bit wasm span descriptor to Go memory slice
func asMemorySlice(context wasmer.InstanceContext, span C.int64_t) []byte {
	memory := context.Memory().Data()
	ptr, size := utils.Int64ToOffsetAndSize(int64(span))
	return memory[ptr : ptr+size]
}

// importsNodeRuntime returns the WASM imports for the node runtime.
func importsNodeRuntime() (imports *wasmer.Imports, err error) {
	imports = wasmer.NewImports()
	// Note imports are closed by the call to wasm.Instance.Close()

	for _, toRegister := range []struct {
		importName     string
		implementation interface{}
		cgoPointer     unsafe.Pointer
	}{
		{"ext_allocator_free_version_1", ext_allocator_free_version_1, C.ext_allocator_free_version_1},
		{"ext_allocator_malloc_version_1", ext_allocator_malloc_version_1, C.ext_allocator_malloc_version_1},
		{"ext_storage_set_version_1", ext_storage_set_version_1, C.ext_storage_set_version_1},
	} {
		_, err = imports.Namespace("env").AppendFunction(toRegister.importName, toRegister.implementation, toRegister.cgoPointer)
		if err != nil {
			return nil, fmt.Errorf("importing function: %w", err)
		}
	}

	return imports, nil
}
