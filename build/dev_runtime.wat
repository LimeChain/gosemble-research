(module
  (type $t0 (func (param i32) (result i32)))
  (type $t1 (func (param i32)))
  (type $t2 (func (result i32)))
  (type $t3 (func (param i32 i32)))
  (type $t4 (func))
  (type $t5 (func (param i32 i32 i32)))
  (type $t6 (func (param i32 i32) (result i64)))
  (import "env" "memory" (memory $env.memory 2))
  (import "env" "ext_allocator_malloc_version_1" (func $ext_allocator_malloc_version_1 (type $t0)))
  (import "env" "ext_allocator_free_version_1" (func $ext_allocator_free_version_1 (type $t1)))
  (func $runtime.alloc (type $t2) (result i32)
    (local $l0 i32) (local $l1 i32) (local $l2 i32) (local $l3 i32)
    global.get $g0
    i32.const 16
    i32.sub
    local.tee $l0
    global.set $g0
    i32.const 0
    local.set $l1
    loop $L0 (result i32)
      block $B1
        block $B2
          block $B3
            i32.const 0
            i32.load offset=1032
            local.tee $l2
            i32.const -17
            i32.ge_u
            br_if $B3
            block $B4
              local.get $l2
              i32.const 17
              i32.add
              local.tee $l3
              i32.const 0
              i32.load offset=1024
              local.tee $l2
              i32.le_u
              br_if $B4
              local.get $l1
              i32.const 1
              i32.and
              i32.eqz
              br_if $B1
              block $B5
                loop $L6
                  local.get $l2
                  i32.eqz
                  br_if $B5
                  local.get $l3
                  local.get $l2
                  i32.le_u
                  br_if $B4
                  i32.const 0
                  local.get $l2
                  i32.const 1
                  i32.shl
                  local.tee $l2
                  i32.store offset=1024
                  br $L6
                end
              end
              i32.const 0
              i32.const -1
              i32.store offset=1024
            end
            i32.const 17
            call $ext_allocator_malloc_version_1
            local.tee $l2
            br_if $B2
            local.get $l1
            i32.const 1
            i32.and
            i32.eqz
            br_if $B1
          end
          unreachable
          unreachable
        end
        local.get $l2
        i64.const 4294967296
        i64.store offset=8 align=4
        local.get $l2
        i64.const 0
        i64.store align=4
        local.get $l0
        i32.const 8
        i32.add
        i64.const 0
        i64.store
        local.get $l0
        i64.const 0
        i64.store
        i32.const 1040
        local.get $l2
        call $_*runtime.memTreap_.insert
        local.get $l2
        i32.const 0
        i32.store8 offset=16
        i32.const 0
        i32.const 0
        i32.load offset=1032
        i32.const 17
        i32.add
        i32.store offset=1032
        local.get $l0
        i32.const 16
        i32.add
        global.set $g0
        local.get $l2
        i32.const 16
        i32.add
        return
      end
      call $runtime.GC
      i32.const 1
      local.set $l1
      br $L0
    end)
  (func $_*runtime.memTreap_.insert (type $t3) (param $p0 i32) (param $p1 i32)
    (local $l2 i32) (local $l3 i32)
    block $B0
      block $B1
        local.get $p0
        i32.eqz
        br_if $B1
        block $B2
          local.get $p0
          i32.load
          br_if $B2
          local.get $p0
          local.get $p1
          i32.store
          br $B0
        end
        local.get $p1
        i32.eqz
        br_if $B1
        local.get $p0
        local.set $l2
        loop $L3
          local.get $p1
          local.get $l2
          i32.load
          i32.store
          local.get $p1
          call $_*runtime.memTreapNode_.parentSlot
          i32.load
          local.set $l3
          local.get $p1
          call $_*runtime.memTreapNode_.parentSlot
          local.set $l2
          local.get $l3
          br_if $L3
        end
        local.get $l2
        local.get $p1
        i32.store
        local.get $p1
        call $_*runtime.memTreapNode_.priority
        local.set $l2
        loop $L4
          local.get $p1
          i32.load
          i32.eqz
          br_if $B0
          local.get $l2
          local.get $p1
          i32.load
          call $_*runtime.memTreapNode_.priority
          i32.le_u
          br_if $B0
          local.get $p0
          local.get $p1
          local.get $p1
          i32.load
          call $_*runtime.memTreap_.rotate
          br $L4
        end
      end
      unreachable
      unreachable
    end)
  (func $runtime.GC (type $t4)
    (local $l0 i32) (local $l1 i32) (local $l2 i32) (local $l3 i32)
    global.get $g0
    i32.const 16
    i32.sub
    local.tee $l0
    global.set $g0
    block $B0
      i32.const 0
      i32.load offset=1040
      i32.eqz
      br_if $B0
      i32.const 1040
      local.set $l1
      block $B1
        loop $L2
          local.get $l1
          i32.load
          local.tee $l2
          i32.eqz
          br_if $B1
          local.get $l2
          i32.const 8
          i32.add
          local.set $l1
          local.get $l2
          i32.load offset=8
          br_if $L2
        end
        local.get $l2
        i32.const 16
        i32.add
        local.set $l3
        i32.const 1040
        local.set $l1
        loop $L3
          local.get $l1
          i32.load
          local.tee $l2
          i32.eqz
          br_if $B1
          local.get $l2
          i32.const 4
          i32.add
          local.set $l1
          local.get $l2
          i32.load offset=4
          br_if $L3
        end
        i32.const 0
        local.get $l3
        i32.store offset=1044
        i32.const 0
        local.get $l2
        local.get $l2
        i32.load offset=12
        i32.add
        i32.const 16
        i32.add
        i32.store offset=1048
        i32.const 1024
        i32.const 66608
        call $runtime.scan
        i32.const 1068
        local.set $l2
        loop $L4
          block $B5
            local.get $l2
            i32.load
            local.tee $l2
            br_if $B5
            loop $L6
              block $B7
                i32.const 0
                i32.load offset=1056
                br_if $B7
                i32.const 0
                i32.load offset=1040
                local.set $l2
                loop $L8
                  block $B9
                    block $B10
                      local.get $l2
                      i32.eqz
                      br_if $B10
                      block $B11
                        local.get $l2
                        i32.load offset=4
                        i32.eqz
                        br_if $B11
                        local.get $l2
                        i32.load offset=4
                        local.set $l2
                        br $L8
                      end
                      block $B12
                        local.get $l2
                        i32.load offset=8
                        i32.eqz
                        br_if $B12
                        local.get $l2
                        i32.load offset=8
                        local.set $l2
                        br $L8
                      end
                      block $B13
                        local.get $l2
                        i32.load
                        local.tee $l1
                        i32.eqz
                        br_if $B13
                        local.get $l2
                        call $_*runtime.memTreapNode_.parentSlot
                        i32.const 0
                        i32.store
                        br $B9
                      end
                      i32.const 0
                      i32.const 0
                      i32.store offset=1040
                      br $B9
                    end
                    i32.const 0
                    i32.const 0
                    i32.load offset=1064
                    i32.store offset=1040
                    local.get $l0
                    i32.const 0
                    i32.store offset=8
                    i32.const 0
                    i32.const 0
                    i32.store offset=1064
                    br $B0
                  end
                  i32.const 0
                  i32.const 0
                  i32.load offset=1032
                  local.get $l2
                  i32.load offset=12
                  i32.sub
                  i32.const -16
                  i32.add
                  i32.store offset=1032
                  local.get $l2
                  call $ext_allocator_free_version_1
                  local.get $l1
                  local.set $l2
                  br $L8
                end
              end
              i32.const 0
              i32.load offset=1056
              local.tee $l2
              i32.eqz
              br_if $B1
              i32.const 0
              local.get $l2
              i32.load offset=4
              local.tee $l1
              i32.store offset=1056
              block $B14
                local.get $l1
                br_if $B14
                i32.const 0
                i32.const 0
                i32.store offset=1052
              end
              local.get $l2
              i64.const 0
              i64.store offset=4 align=4
              local.get $l2
              i32.const 16
              i32.add
              local.tee $l1
              local.get $l1
              local.get $l2
              i32.load offset=12
              i32.add
              call $runtime.scan
              i32.const 1064
              local.get $l2
              call $_*runtime.memTreap_.insert
              br $L6
            end
          end
          local.get $l2
          i32.const 8
          i32.add
          local.tee $l1
          local.get $l1
          local.get $l2
          i32.load offset=4
          i32.const 2
          i32.shl
          i32.add
          call $runtime.scan
          br $L4
        end
      end
      unreachable
      unreachable
    end
    local.get $l0
    i32.const 16
    i32.add
    global.set $g0)
  (func $runtime.scan (type $t3) (param $p0 i32) (param $p1 i32)
    (local $l2 i32) (local $l3 i32) (local $l4 i32)
    local.get $p1
    i32.const -4
    i32.and
    local.set $l2
    local.get $p0
    i32.const 3
    i32.add
    i32.const -4
    i32.and
    local.set $l3
    block $B0
      block $B1
        loop $L2
          local.get $l3
          local.get $l2
          i32.ge_u
          br_if $B1
          block $B3
            i32.const 0
            i32.load offset=1044
            local.get $l3
            i32.load
            local.tee $p1
            i32.gt_u
            br_if $B3
            i32.const 0
            i32.load offset=1048
            local.get $p1
            i32.lt_u
            br_if $B3
            i32.const 1040
            local.set $p0
            loop $L4
              local.get $p0
              i32.load
              local.tee $p0
              i32.eqz
              br_if $B3
              block $B5
                block $B6
                  local.get $p1
                  local.get $p0
                  i32.const 16
                  i32.add
                  local.tee $l4
                  i32.lt_u
                  br_if $B6
                  local.get $p0
                  i32.load offset=12
                  local.get $l4
                  i32.add
                  local.get $p1
                  i32.gt_u
                  br_if $B5
                end
                local.get $p0
                i32.const 4
                i32.const 8
                local.get $p1
                local.get $p0
                i32.gt_u
                select
                i32.add
                local.set $p0
                br $L4
              end
            end
            block $B7
              loop $L8
                block $B9
                  local.get $p0
                  i32.load offset=4
                  br_if $B9
                  local.get $p0
                  i32.load offset=8
                  br_if $B9
                  local.get $p0
                  i32.load
                  br_if $B9
                  i32.const 0
                  i32.const 0
                  i32.store offset=1040
                  br $B7
                end
                block $B10
                  local.get $p0
                  i32.load offset=4
                  br_if $B10
                  local.get $p0
                  i32.load offset=8
                  br_if $B10
                  local.get $p0
                  call $_*runtime.memTreapNode_.parentSlot
                  i32.const 0
                  i32.store
                  br $B7
                end
                block $B11
                  local.get $p0
                  i32.load offset=4
                  i32.eqz
                  br_if $B11
                  local.get $p0
                  i32.load offset=8
                  br_if $B11
                  local.get $p0
                  local.get $p0
                  i32.load offset=4
                  call $_*runtime.memTreap_.replace
                  br $B7
                end
                block $B12
                  block $B13
                    local.get $p0
                    i32.load offset=8
                    i32.eqz
                    br_if $B13
                    local.get $p0
                    i32.load offset=4
                    i32.eqz
                    br_if $B12
                  end
                  i32.const 1040
                  local.get $p0
                  i32.const 4
                  i32.const 8
                  local.get $p0
                  i32.load offset=4
                  call $_*runtime.memTreapNode_.priority
                  local.get $p0
                  i32.load offset=8
                  call $_*runtime.memTreapNode_.priority
                  i32.gt_u
                  select
                  i32.add
                  i32.load
                  local.get $p0
                  call $_*runtime.memTreap_.rotate
                  br $L8
                end
              end
              local.get $p0
              local.get $p0
              i32.load offset=8
              call $_*runtime.memTreap_.replace
            end
            local.get $p0
            i64.const 0
            i64.store align=4
            local.get $p0
            i32.const 0
            i32.store offset=8
            block $B14
              block $B15
                i32.const 0
                i32.load offset=1052
                br_if $B15
                i32.const 1056
                local.set $p1
                br $B14
              end
              i32.const 0
              i32.load offset=1052
              local.tee $p1
              i32.eqz
              br_if $B0
              local.get $p1
              i32.const 4
              i32.add
              local.set $p1
            end
            local.get $p1
            local.get $p0
            i32.store
            i32.const 0
            i32.load offset=1052
            local.set $p1
            i32.const 0
            local.get $p0
            i32.store offset=1052
            local.get $p0
            local.get $p1
            i32.store offset=8
          end
          local.get $l3
          i32.const 4
          i32.add
          local.set $l3
          br $L2
        end
      end
      return
    end
    unreachable
    unreachable)
  (func $_*runtime.memTreapNode_.parentSlot (type $t0) (param $p0 i32) (result i32)
    (local $l1 i32)
    block $B0
      local.get $p0
      i32.eqz
      br_if $B0
      block $B1
        local.get $p0
        i32.load
        local.tee $l1
        local.get $p0
        i32.ge_u
        br_if $B1
        local.get $l1
        i32.eqz
        br_if $B0
        local.get $l1
        i32.const 4
        i32.add
        return
      end
      local.get $l1
      i32.eqz
      br_if $B0
      local.get $l1
      i32.const 8
      i32.add
      return
    end
    unreachable
    unreachable)
  (func $_*runtime.memTreapNode_.priority (type $t0) (param $p0 i32) (result i32)
    local.get $p0
    i32.const -1640531527
    i32.mul)
  (func $_*runtime.memTreap_.rotate (type $t5) (param $p0 i32) (param $p1 i32) (param $p2 i32)
    (local $l3 i32)
    block $B0
      block $B1
        block $B2
          block $B3
            local.get $p1
            local.get $p2
            i32.le_u
            br_if $B3
            local.get $p0
            i32.eqz
            br_if $B0
            local.get $p0
            i32.load
            local.get $p2
            i32.eq
            br_if $B2
            local.get $p2
            call $_*runtime.memTreapNode_.parentSlot
            local.set $p0
            br $B2
          end
          local.get $p0
          i32.eqz
          br_if $B0
          block $B4
            local.get $p0
            i32.load
            local.get $p2
            i32.eq
            br_if $B4
            local.get $p2
            call $_*runtime.memTreapNode_.parentSlot
            local.set $p0
          end
          local.get $p0
          local.get $p1
          i32.store
          local.get $p1
          i32.eqz
          br_if $B0
          local.get $p2
          i32.eqz
          br_if $B0
          local.get $p1
          i32.load offset=4
          local.set $p0
          local.get $p2
          i32.load
          local.set $l3
          local.get $p1
          local.get $p2
          i32.store offset=4
          local.get $p1
          local.get $l3
          i32.store
          local.get $p2
          local.get $p0
          i32.store offset=8
          br $B1
        end
        local.get $p0
        local.get $p1
        i32.store
        local.get $p2
        i32.eqz
        br_if $B0
        local.get $p1
        i32.load offset=8
        local.set $p0
        local.get $p2
        i32.load
        local.set $l3
        local.get $p1
        local.get $p2
        i32.store offset=8
        local.get $p1
        local.get $l3
        i32.store
        local.get $p2
        local.get $p0
        i32.store offset=4
      end
      local.get $p2
      local.get $p1
      i32.store
      block $B5
        local.get $p0
        i32.eqz
        br_if $B5
        local.get $p0
        local.get $p2
        i32.store
      end
      return
    end
    unreachable
    unreachable)
  (func $_*runtime.memTreap_.replace (type $t3) (param $p0 i32) (param $p1 i32)
    block $B0
      local.get $p1
      i32.eqz
      br_if $B0
      local.get $p0
      i32.eqz
      br_if $B0
      local.get $p1
      local.get $p0
      i32.load
      i32.store
      block $B1
        i32.const 0
        i32.load offset=1040
        local.get $p0
        i32.ne
        br_if $B1
        i32.const 0
        local.get $p1
        i32.store offset=1040
        return
      end
      local.get $p1
      call $_*runtime.memTreapNode_.parentSlot
      local.get $p1
      i32.store
      return
    end
    unreachable
    unreachable)
  (func $Core_version (type $t6) (param $p0 i32) (param $p1 i32) (result i64)
    (local $l2 i32) (local $l3 i32) (local $l4 i32) (local $l5 i32) (local $l6 i32)
    global.get $g0
    i32.const 80
    i32.sub
    local.tee $l2
    global.set $g0
    local.get $l2
    i32.const 52
    i32.add
    local.tee $l3
    i64.const 0
    i64.store align=4
    local.get $l2
    i32.const 60
    i32.add
    local.tee $l4
    i64.const 0
    i64.store align=4
    local.get $l2
    i32.const 0
    i32.store offset=15 align=1
    local.get $l2
    i64.const 0
    i64.store offset=8
    local.get $l2
    i64.const 0
    i64.store offset=68 align=4
    local.get $l2
    i32.const 7
    i32.store offset=44
    i32.const 0
    i32.load offset=1068
    local.set $l5
    i32.const 0
    local.get $l2
    i32.const 40
    i32.add
    i32.store offset=1068
    local.get $l2
    local.get $l5
    i32.store offset=40
    local.get $l2
    i32.const 48
    i32.add
    local.get $l2
    i32.const 8
    i32.add
    i32.store
    local.get $l2
    i32.const 97
    i32.store8 offset=18
    local.get $l2
    i32.const 29793
    i32.store16 offset=16
    local.get $l2
    i32.const 101
    i32.store8 offset=14
    local.get $l2
    i32.const 28009
    i32.store16 offset=12
    local.get $l2
    i32.const 1953396050
    i32.store offset=8
    local.get $l2
    i32.const 68
    i32.store8 offset=15
    local.get $l3
    local.get $l2
    i32.const 24
    i32.add
    i32.store
    local.get $l2
    i32.const 0
    i32.store offset=31 align=1
    local.get $l2
    i64.const 0
    i64.store offset=24
    local.get $l2
    i32.const 97
    i32.store8 offset=34
    local.get $l2
    i32.const 29793
    i32.store16 offset=32
    local.get $l2
    i32.const 68
    i32.store8 offset=31
    local.get $l2
    i32.const 101
    i32.store8 offset=30
    local.get $l2
    i32.const 28009
    i32.store16 offset=28
    local.get $l2
    i32.const 1953396050
    i32.store offset=24
    local.get $l2
    i32.const 64
    i32.add
    local.get $p0
    i32.store
    local.get $l4
    local.get $p0
    i32.store
    local.get $l2
    i32.const 56
    i32.add
    local.get $p0
    i32.store
    block $B0
      block $B1
        local.get $p1
        i32.const 0
        i32.lt_s
        br_if $B1
        local.get $p0
        i32.eqz
        local.get $p1
        i32.const 0
        i32.ne
        i32.and
        br_if $B1
        local.get $p1
        i32.const 11
        local.get $p1
        i32.const 11
        i32.lt_u
        select
        local.set $l3
        local.get $l2
        i32.const 24
        i32.add
        local.set $l4
        local.get $p0
        local.set $l6
        loop $L2
          local.get $l3
          i32.eqz
          br_if $B0
          local.get $l6
          local.get $l4
          i32.load8_u
          i32.store8
          local.get $l6
          i32.const 1
          i32.add
          local.set $l6
          local.get $l4
          i32.const 1
          i32.add
          local.set $l4
          local.get $l3
          i32.const -1
          i32.add
          local.set $l3
          br $L2
        end
      end
      unreachable
      unreachable
    end
    local.get $l2
    i32.const 68
    i32.add
    call $runtime.alloc
    local.tee $l3
    i32.store
    local.get $l2
    i32.const 72
    i32.add
    local.get $l3
    i32.store
    local.get $l3
    i32.const 127
    i32.store8
    i32.const 0
    local.get $l5
    i32.store offset=1068
    local.get $l2
    i32.const 80
    i32.add
    global.set $g0
    local.get $p1
    i64.extend_i32_u
    i64.const 32
    i64.shl
    local.get $p0
    i64.extend_i32_s
    i64.or)
  (table $__indirect_function_table 1 1 funcref)
  (global $g0 (mut i32) (i32.const 66608))
  (global $__heap_base i32 (i32.const 66608))
  (global $__data_end i32 (i32.const 1072))
  (export "__heap_base" (global 1))
  (export "Core_version" (func $Core_version))
  (export "__data_end" (global 2))
  (export "__indirect_function_table" (table 0))
  (data $d0 (i32.const 1024) "@\00\00\00")
  (data $d1 (i32.const 1032) "\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00\00"))
