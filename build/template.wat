(module
  (type $t0 (func (param i32) (result i32)))
  (type $t1 (func (param i32 i32) (result i64)))
  (import "env" "memory" (memory $env.memory 2))
  (import "env" "ext_allocator_malloc_version_1" (func $main.extAllocatorMallocVersion1 (type $t0)))
  (func $Core_version (type $t1) (param $p0 i32) (param $p1 i32) (result i64)
    i32.const 0
    call $main.extAllocatorMallocVersion1
    drop
    i64.const 1134)
  (table $__indirect_function_table 1 1 funcref)
  (global $g0 (mut i32) (i32.const 66560))
  (global $__data_end i32 (i32.const 1024))
  (global $__heap_base i32 (i32.const 66560))
  (export "Core_version" (func $Core_version))
  (export "__data_end" (global 1))
  (export "__heap_base" (global 2))
  (export "__indirect_function_table" (table 0)))
