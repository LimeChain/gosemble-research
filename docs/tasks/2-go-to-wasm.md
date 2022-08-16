# Research Feasibility of a Go Runtime - Oct 2022

## Abstract

The idea of writing Polkadot Runtimes in Go is exciting, mainly because of Go's simplicity and automatic memory management. However Polkadot's design desicions around memory management (managed by the Host) defeat one of Go's main selling points. Although the language specification doesn't mention how it should manage its memory, the Go community recognizes it as a language with GC, and anything else would be more like a different language with a similar syntax. So aren't we setting up ourselves for failure with Go right from the beginning?

Questions we are gonna try to answer:
* What are the design decisions behind Substrate - Webassembly specification, Host/Runtime interaction, Runtime memory management?
* Which Webassembly features are used by Substrate? Any limitations? Any proposals that might help in the future?
* Which compiler toolchain is best suited in the case of producing Wasm with GC runtime - LLVM or Binaryen?
* How Go and TinyGo internals work and differ from each other - compiler, runtime, GC?
* What are the challenges of implementing a Substrate Runtime with GC managed language like Go?
* What are the steps required to start implementing new compiler/runtime with GC and external allocator?


## Substrate design decisions

**WebAssembly specification**

The Wasm runtime module targets [WebAssembly MVP](https://github.com/WebAssembly/design/blob/main/MVP.md) without any extensions enabled and domain-specific API.
* [x] module exports API functions (the business logic).
* [x] module imports Host provided functions (`ext_allocator_free/ext_allocator_malloc`).
* [ ] module imports Host provided memory.
* [ ] module exports linker specific globals (`__heap_base`, `__data_end`).
* [ ] module exports `__indirect_function_table` (it is WIP and not enabled currently).

```
Type: wasm
...
Imports:
  Functions:
    "env"."ext_allocator_malloc_version_1": [I32] -> [I32]
    "env"."ext_allocator_free_version_1": [I32] -> []
    ...
  Memories:
    "env"."memory": not shared (20 pages..)
  Tables:
  Globals:
Exports:
  Functions:
    "Core_version": [I32, I32] -> [I64]
    "Core_execute_block": [I32, I32] -> [I64]
    "Core_initialize_block": [I32, I32] -> [I64]
    ...
  Memories:
  Tables:
    "__indirect_function_table": FuncRef (352..352)
  Globals:
    "__data_end": I32 (constant)
    "__heap_base": I32 (constant)
 _________________________________________________________________________________________
| HOST                                                                                    |
|                                       _________________________________________         |
|                                      |                                         |        |
|                                      | ext_allocator_malloc/ext_allocator_free |        |
|                                      |_________________________________________|        |
|      ______________________________________________________|_______________________     |
|     | WASM                                                 | (imported)            |    |
|     |                                                      |                       |    |
|     |             _________________________________________▼_____________________  |    |
|     | (imported) |          |                  |           [-----xxx]            | |    |
|  MEMORY  -----►  | Data     |         ◄- Stack | Heap -►        [xxx-----]       | |    |
|     |            |__________|__________________|_________________________________| |    |
|     |            0     __data_end         __heap_base           ▲      max memory  |    |
|     |                                                           |                  |    |
|     |                                                           GC                 |    |
|     |______________________________________________________________________________|    |
|                                                                                         |
|_________________________________________________________________________________________|

```

**WASI interface requirements**

Polkadot is non-browser environment, but it is not an OS and doesn't seeks to provide a system-level API comparable to an OS like files, networking, or a major part of other things provided by WASI.

**Imported or exported memory**

Working with declared exported memory is almost certainly still supported, but does not matter too much. In fact, this is how it worked in the beginning and importing memory. Imported memory works a little bit better than exported memory since it avoids some edge cases, although it also has some downsides.

**Exported globals**

The beginning of this heap is marked by the `__heap_base` symbol exported by the Wasm module. No memory should be allocated below that address, to avoid clashes with the stack and data section.

**SCALE codec**

Runtimes only understand byte code and SCALE helps applications efficiently decode and encode data for them. Runtime data needs to be as light as possible, a capability that SCALE provides. And its performance is owed to it being built for LE architectures, which is compatible with Wasm environments.

**Runtime function calls**

Each call into the runtime is done with fresh memory, allocated via the host allocator, allocations do not persist between calls. The runtime uses the same host provided allocator for all allocations, so the host is in charge of that. Data passing the Runtime API are always scale encoded. Host API call on the other hand try to avoid all encoding.

**External memory management**

The design in which allocation functions are on the host side is dictated by the fact that some of the host functions might return buffers of data of unknown size. That means that the wasm code cannot efficiently provide buffers upfront. For example, let's examine the host function that returns a given storage value. The storage value's size is not known upfront in the general case, so the wasm caller cannot pre-allocate the buffer upfront. A potential solution is to first call the host function without a buffer, which will return the value's size, and then do the second call passing a buffer of the required size. For some host functions, caches could be put in place for mitigation, some other functions cannot be implemented in the such model at all. To solve this problem, it was chosen to place the allocator on the host side.

Note, however, that this is not the only possible solution. For instance, there is an ongoing discussion about moving the allocator into the wasm: [1](https://github.com/paritytech/substrate/issues/11883)

Notably, the allocator maintains some of its data structures inside the linear memory and some other structures outside.

Before each function call to the Runtime, some memory space is reserved via the allocator, either for sahring input data or results. After the call that space is reclaimed via the allocator.

**Language with an automatic memory management**

Theoretically, it might be possible, but the support would be limited, performance might be unsatisfactory and the toolchain would need to polyfill the GC. To support an automatic memory management, the [GC proposal](https://github.com/WebAssembly/gc/blob/main/proposals/gc/Overview.md) might be handy. But the Wasm runtime supports only WebAssembly MVP currently, also the GC proposal is under development and it is not yet clear if Polkadot will be able to leverage the GC proposal. Potential problems include determinism (is there anything in GC that causes ND? Can it be tamed efficiently?) and safety (Is it possible for a host to limit the resource consumption reliably and deterministically?).

**Support of concurrency**

The runtime executes in serial, the parallelism is accomplished through a network of parallel running chains (Parachains). It is primarily because creating a semantic for deterministic parallel executing is really difficult in general.


## WebAssembly MVP features, limitations and future proposals

**Features**
* instruction format for bytecode stack machine
* fast execution
* compact
* portable

* Little-endian byte order when translating between values and bytes.
* Linear memory, a contiguous, byte-addressable, linear address space, spanning from offset 0 and extending up to a varying memory size. Can be resized, but it can only be grown. (`grow_memory`, `memory.grow`)
* Single specially-designated default linear memory which is the linear memory accessed by all the memory operators.
* Linear memory cannot be shared between threads of execution. 

**Limitations**

Linear memories (default or otherwise) can either be imported or defined inside the module. After import or definition, there is no difference when accessing a linear memory whether it was imported or defined internally.
In the MVP, linear memory cannot be shared between threads of execution. The addition of threads 🦄 will allow this.

**Proposals**
* Allow setting protection and creating mappings within the contiguous linear memory.
* In the MVP, there are only default linear memories but new memory operators 🦄 may be added after the MVP which can also access non-default memories.
* WebAssembly/gc proposal most likely won't help [1](https://github.com/WebAssembly/gc/issues/59).
* there is no WebAssembly specification for exports and runtime behavior around allocation.


## LLVM vs Binaryen compiler backends

* LLVM is much more powerful as an optimizer.
* LLVM does not support Wasm GC, and the future there is unclear. In general, GC is not a main focus for LLVM (almost all the languages using it use linear memory, C, C++, Rust, Swift, Zig, etc.).
* LLVM supports wasm object files, DWARF, and other things which are very useful in the non-GC world (they may also help in GC in the future, that's unclear; we'll add support to Binaryen as needed).

* Binaryen compiles much more quickly.
* Binaryen is much smaller.
* Binaryen is a good choice for languages with GC and intend to compile to Wasm GC. And we will be able to compile Wasm GC to Wasm MVP, as a polyfill until Wasm GC is everywhere, but not the opposite.


## Go

### Compiler

The default compiler is `gc`. There are also `gccgo` which uses the GCC back-end and `gollvm` which uses the LLVM infrastructure (somewhat less mature).

```
// Frontend Compiler (lexing, parsing, typechecking)

CODE -> tokenizer/lexer/scanner (lexical analysis) -> TOKENS STREAM
TOKENS STREAM -> go/parser (syntactic analysis) -> AST (annotated)
AST -> go/types (type check) -> AST (with types)
AST (with types) -> golang.org/x/tools/go/ssa (convert) -> SSA

// Optimization

* variable capturing, inlining, escape analysis, closure rewriting, walk

SSA (higher level with Go-specific constructs like interfaces and goroutines) -> convert (tinygo) -> LLVM IR
LLVM IR -> optimize (llvm) -> LLVM IR (optimize by a mixture of handpicked LLVM optimization passes, TinyGo-specific optimizations ()

// Backend Compiler

LLVM IR (optimized) -> convert (llvm) -> MACHINE CODE
MACHINE CODE -> convert (llvm) -> OBJECT FILE
OBJECT FILE -> link (llvm) -> EXECUTABLE
```

### Runtime

Implements GC, scheduler included in every Go program. Contains a lot of type information at runtime.

**Memory Management**

Process
* does not have direct access to the physical memory
* virtual memory abstracts the access to the physical memory (via segmentation and page tables)

Process Memory Layout
```
 ______________________
|        STACK         | Function Stack Frames
|----------------------| ◄--- Stack Pointer
|          |           |
|          ▼           |
|                      |
|          ▲           |
|          |           |
|----------------------|
|         HEAP         | Dynamic Allocated Variables
|----------------------|
|         BSS          | Uninitializaed Static Variables
|----------------------|
|         DATA         | Initializaed Static Variables 
|----------------------|
|         TEXT         | Code
|______________________|
```

Stack
* managed by the compiler
* elastic
* one stack per goroutine

Heap
* allocated by the memory allocator and collected by the garbage collector
* it is not an entity and there is no linear containment of memory that defines the Heap
* any memory reserved for application use in the process space is available for heap memory allocation

Go uses escape analysis and Garbage Collector. Allocator is tightly coupled with the Garbage Collector.
```
   Mutator (Process)
      |
      | malloc/free
      |
      ▼
  Allocator
      |
      | syscall mmap/munmap (address, size, permissions, flags)
      |
 _____▼____
|          | 
|   Heap   | ◄------ Collector
|__________|
```

Allocator
* allocate new blocks with the correct size
* deals with fragmentation (merge smaller block to allow allocation of larger ones)



```
                                         BLOCK OF MEMORY
 ____________________________________________________________________________________________
|    HEADER (x bits)   |          PAYLOAD          |     PADDING      |    FOOTER (x bits)   |
|                      |                           |                  |                      |
|                      |       (if allocated)      |    (optional)    |                      |
| block size | 0 0 0/1 |                           |                  | block size | 0 0 0/1 |
|______________________|___________________________|__________________|______________________|
                       ▲
                       |
                      ptr

(x = 0:free | 1:allocated)

Double linked-list of free objects using the block header/footer.
 ______________________________________________________________
|________|________|________|________|________|________|________|
  64 bits  64 bits  64 bits  64 bits  64 bits  64 bits  64 bits (aligned by 8 bytes)


Bump Allocator (needs Garbage Collector to reclaim memory)
 _______________________________________________________________
|        |                    |        |                        |
| Object |        Free        | Object |          Free          |
|________|____________________|________|________________________|
                                       ▲

Free List Allocator
 _______________________________________________________________
|        |                    |        |                        |
| Object |        Free        | Object |          Free          |
|________|____________________|________|________________________|
                    |                  ▲            ▲
                    |_______________________________|
```

Garbage Collector
* tracking memory allocations in heap memory
* releasing allocations that are no longer needed
* keeping allocatiosn that are still in use

* tracing garbage collector (not reference counting)
* hybrid stop-the-world/concurrent collector (stop-the-world part limited by a 10ms deadline)
* CPU cores dedicated to running the concurrent collector
* tri-color mark-and-sweep algorithm

* non-generational
* non-compacting
* fully precise
* incurs a small cost if the program is moving pointers around
* lower latency, but most likely also lower throughput

*Collection*
1. Mark Setup (stop the world)
  * turn on write barrier
  * stop all goroutines

2. Marking (concurrent)
  * inspect the stack to find root pointers to the heap
  * traverse the heap graph from those root pointers
  * mark values on the heap that are still in use
  * slow down allocations to speed up collection

3. Mark Termination (stop the world)
  * turn the write barrier off
  * various cleaup tasks
  * next collection goal is acalulated

*Sweeping*
Freeing Heap Memory
* occurs when the goroutines attempt to allocate new heap memory
* the latency of sweeping is added to the cost of performing new allocation

Compiler decides (via escape analysis) when a value should be allocated on the Heap
* sharing down (passign pointers) typically stays on the Stack
* sharing up (returning pointers) typically escapes on the Heap

* when the value could be referenced after the function that constructed it returns
* if the value is too large to fit on the stack
* when the compiler doesn't know the size of the value at compile time


RSS MEMORY (stack + heap + globals)

```
			GOROUTINE (main)           GOROUTINE (1)
					 |                         |
					 |                         |
					 ▼                         ▼
				 STACK                     STACK (1)
 ______________________
|                      |
|                      |
|                      |
|                      |
|   	main (FRAME)     |
|  __________________  |
| |                  | |
| |                  | |
| |                  | |
| |  newObj (FRAME)  | |
| |  ______________  | |
| | |              | | |
| | |              | | |
| | |              | | |
| | |              | | |
| | |              | | |
| | |              | | |
| | |              | | |
| | |______________| | |
| |                  | |
| |         o        | |
| |_________|________| |
|___________|__________|
						|
 	    ______| 				   HEAP
     |
  ___|_________________________________________________
 |   |                                                 |
 |   ▼                                                 |
 |  obj [--------]                                     |
 |                                                     |
 |                                                     |
 |_____________________________________________________|
start                                                 end

push main
push o
push newObj
--► obj heap
pop newObj
pop main

--------------------------------------------

In-Memory Representation of Go Types

byte (uint8)
rune (int32)
bool (int8)
int
float
complex

string

array
struct

slice

map

channel
pointer
function

interfaces
```

**Concurrency**

The scheduler runs goroutines, pauses and resumes them on blocking channel ops. or mutex ops, coordinates blocking system calls, io, runtime GC. Goroutines are use space threads managed by the runtime.

**Parallelism**


## TinyGo

It is a subset of Go with very different goals from the standard Go. It is a new compiler and runtime aimed to support many different small embedded devices with a single processor core that require certain optimizations mostly toward size.

*Goals*
* Have very small binary sizes.
* Support for most common microcontroller boards.
* Be usable on the web using WebAssembly.
* Good CGo support, with no more overhead than a regular function call.
* Support most standard library packages and compile most Go code without modification.

*Non-goals*
* Using more than one core.
* Be efficient while using zillions of goroutines. However, good goroutine support is certainly a goal.
* Be as fast as `gc`. However, LLVM will probably be better at optimizing certain things so TinyGo might actually turn out to be faster for number crunching.
* Be able to compile every Go program out there.

### Compiler

The compiler uses (mostly) the standard library to parse Go programs and LLVM to optimize the code and generate machine code for the target architecture.

**Pipeline (Lowering)**
```
// Frontend Compiler (lexing, parsing, typechecking)

CODE -> parse (go/parser) -> AST
AST -> type check (go/types) -> AST (with types)
AST (with types) -> convert (golang.org/x/tools/go/ssa) -> SSA

// Optimization

SSA (higher level with Go-specific constructs like interfaces and goroutines) -> convert (tinygo) -> LLVM IR
LLVM IR -> optimize (llvm) -> LLVM IR (optimize by a mixture of handpicked LLVM optimization passes, TinyGo-specific optimizations (escape analysis, string-to-[]byte optimizations, etc.) and custom lowering.)

// Backend Compiler

LLVM IR (optimized) -> convert (llvm) -> MACHINE CODE
MACHINE CODE -> convert (llvm) -> OBJECT FILE
OBJECT FILE -> link (llvm) -> EXECUTABLE
```

* root: contains the command line interface for the tinygo command and all its subcommands
* builder: orchestrates the build
* loader: loads and typechecks the code, and produces an AST
* compiler: the compiler itself, makes little attempt at optimizing code
* interp: tries to run package initializers at compile time as far as possible
* transform: implements various optimizations necessary to produce working and efficient code

### Runtime

The runtime is written from scratch, optimized for size instead of speed and re-implements some compiler intrinsics and packages:

* Startup code (device specific runtime initialization of memory, timers)
* GC (copied from micropython, super simple, optimized for size, runs the GC once it runs out of memory)
* Memory Allocator
* Goroutines scheduler
* Channels
* Time handling (every chip has a clock)
* Hashmap implementation (optimized for size instead of speed)
* Runtime functions like maps, slice append, defer, ... (various things that are not implemented in the compiler)
* Operations on strings
* `sync` & `reflect` package (strongly connected to the runtime)

**Basic features**

All basic types, slices, all regular control flow including switch, closures and bound methods are supported, `defer` keyword is almost entirely supported, with the exception of deferring some builtin functions, interfaces are quite stable and should work well in almost all cases. Type switches and type asserts are also supported, as well as calling methods on interfaces. The only exception is comparing two interface values.
Maps are usable but not complete. You can use any type as a value, but only some types are acceptable as map keys (strings, integers, pointers, and structs/arrays that contain only these types). Also, they have not been optimized for performance and will cause linear lookup times in some cases.

**Memory Management**

 there's no concurrent GC that would free it meanwhile.

Heap with a garbage collector. While not directly a language feature, garbage collection is important for most Go programs to make sure their memory usage stays in reasonable bounds.

Garbage collection is currently supported on all platforms, although it works best on 32-bit chips. A simple conservative mark-sweep collector is used that will trigger a collection cycle when the heap runs out (that is fixed at compile time) or when requested manually using `runtime.GC()`. Some other collector designs are used for other targets, TinyGo will automatically pick a good GC for a given target.

Careful design may avoid memory allocations in main loops. You may want to compile with `-print-allocs=.` to find out where allocations happen and why they happen.

**Concurrency**

Goroutines and channels work for the most part, for platforms such as WebAssembly the support is a bit more limited (calling a blocking function may for example allocate heap memory).

**Parallelism**

Single core only.


## Technical challenges

**Toolchain support for Wasm for non-browser environments**
* The official Go compiler does not support Wasm for non-browser environments [1](https://github.com/golang/go/issues/31105), [2](https://substrate.stackexchange.com/questions/60/what-is-gossamer-and-how-does-it-compare-to-substrate/89#89), only Wasm with browser specific API. The only options that supports that is TinyGo. 

**Wasm features that are not part of the MVP**
* TinyGo makes use of bulk memory operations which are not supported in the targeted Wasm MVP, such as bulk memory operations (`memory.copy`, `memory.fill`) used to reduce the code size.

**Standard library support**
* The standard library relies on the `reflect` package (most common types like numbers, strings, and structs are supported), which is not fully supported by TinyGo [1](https://github.com/tinygo-org/tinygo/pull/2640). The core primitives and SCALE serialization logic that we intended to reuse from [gossamer](https://github.com/ChainSafe/gossamer) all rely on the `reflect` package.
* Many features of Cgo are still unsupported (#cgo statements are only partially supported).

**External memory allocator**
* According to the Polkadot specification, the Wasm module does not include a memory allocator, it imports memory from the Host and relies on Host imported functions for all heap allocations (`ext_allocator_malloc/ext_allocator_free`). TinyGo implements simple GC and manages its memory by itself, contrary to specification. So it can't work directly on systems where the host wants to manage the memory [1](https://github.com/golang/go/issues/13761).

**Linker globals**
* The runtime is expected to expose `__heap_base` global [1](https://github.com/tinygo-org/tinygo/issues/2045), but TinyGo doesn't support that out of the box, however it is relatively simple to add support for it.

**Developer experience**
* Implementing Wasm functionality makes you go pretty low level and use some "unsafe" constructs.


## Proof of concept (alternative compiler + runtime)

TinyGo design decisions are based on optimizations around small embedded devices, which might not always be suitable or required in Wasm. Also, it will be difficult to support custom and non-standard APIs in the long run, the TinyGo core contributors are a bit against supporting things that are not standardized as is the case with Polkadot Wasm's Runtime specification.
It will be best to implement a solution similar to TinyGo (compiler + runtime + GC with external memory allocator) only targeting Wasm. The frontend-compiler should be mostly the same as in TinyGo and the runtime might not need to be super size-optimized.

**Add Toolchain support for Substrate's Wasm**
1. [x] Fork [TinyGo](https://github.com/radkomih/tinygo).
2. [x] Add build script and separate Dockerfile for building TinyGo with prebuild LLVM (for faster builds).
3. Add new target similar to Rust's `wasm32-unknown-unknown`.
  * [x] add a new directive `polkawasm` to separate the new target from the existing WASM/WASI functionality.
  * [x] remove the `_start` export.
  * [x] remove any JS/WASI exports.
  * [x] remove the memory allocation exports.
  * [x] remove any bulk memory operations and other features that are not part of the Wasm MVP spec. 
  * [x] add compiler flag for exporting `__heap_base`, `__data_end` globals.
  * [x] add compiler flag for declaring the memory as imported.
  * [x] add custom GC implementation that can work with external memory allocator.
  * [x] override the allocation functions used in the GC with such provided by the host.
  * [ ] need better abstractions since the runtime implementation depeneds on third party API that might change in the future.

**Build Wasm Runtime**

1. [ ] Implement the SCALE codec without reflection.
2. Implement the minimal Runtime API (core API).
  * [ ] `Core_version`
  * [ ] `Core_execute_block`
  * [ ] `Core_initialize_block`
3. [x] Utilise our TinyGo fork to compile the `Go` implementation of Polkadot's runtime logic as Wasm module.

**Setup Wasm Host**

1. [x] Setup host runtime environment to execute the compiled Wasm module inside (targeting MVP).
2. [x] Import the host provided functions used inside the Wasm module.
3. [x] Import the host provided memory inside the Wasm module.
4. [x] Read/write from/to the host/Wasm's shared memory.
5. [x] Add implementation of bump allocator.

* [ ] fix the error when passing 0 offset to a Wasm function.
* [ ] in the extracted vm helper, some struct fields are undefined

6. Add tests for correctness and performance.


## Testing guide


