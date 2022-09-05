package main

// /*
// "//export" on a func is actually an import in TinyGo.
// The function needs to be referenced somewhere to be actually exported.
// */

// /*
// 	Allocator: Provides functionality for calling into the memory allocator.
// */

// //go:wasm-module env
// //export ext_allocator_malloc_version_1
// func ext_allocator_malloc_version_1(size int32) int32

// //go:wasm-module env
// //export ext_allocator_free_version_1
// func ext_allocator_free_version_1(ptr int32)

// // /*
// // 	Crypto: Interfaces for working with crypto related types from within the runtime.
// // */
// // func ext_crypto_ed25519_generate_version_1(key_type_id int32, seed int64) int32
// // func ext_crypto_ed25519_verify_version_1(sig int32, msg int64, key int32) int32
// // func ext_crypto_secp256k1_ecdsa_recover_compressed_version_1(sig int32, msg int32) int64
// // func ext_crypto_secp256k1_ecdsa_recover_version_1(sig int32, msg int32) int64
// // func ext_crypto_sr25519_generate_version_1(key_type_id int32, seed int64) int32
// // func ext_crypto_sr25519_public_keys_version_1(key_type_id int64) int64
// // func ext_crypto_sr25519_sign_version_1(key_type_id int32, key int32, msg int64) int64
// // func ext_crypto_sr25519_verify_version_1(sig int32, msg int64, key int32) int32
// // func ext_crypto_sr25519_verify_version_2(sig int32, msg int64, key int32) int32

// /*
// 	Hashing: Interface that provides functions for hashing with diﬀerent algorithms.
// */
// func ext_hashing_blake2_128_version_1(data int64) int32
// func ext_hashing_blake2_256_version_1(data int64) int32
// func ext_hashing_keccak_256_version_1(data int64) int32
// func ext_hashing_twox_128_version_1(data int64) int32
// func ext_hashing_twox_64_version_1(data int64) int32

// /*
// 	Log: Request to print a log message on the host. Note that this will be
// 	only displayed if the host is enabled to display log messages with given level and target.
// */

// //go:wasm-module env
// //export ext_logging_log_version_1
// func ext_logging_log_version_1(level int32, target int64, message int64)

// /*
// 	Miscellaneous: Interface that provides miscellaneous functions for communicating between the runtime and the node.
// */
// func ext_misc_print_hex_version_1(data int64)
// func ext_misc_print_num_version_1(value int64)
// func ext_misc_print_utf8_version_1(data int64)
// func ext_misc_runtime_version_version_1(data int64) int64

// /*
// 	Storage: Interface for manipulating the storage from within the runtime.
// */

// //go:wasm-module env
// //export ext_storage_changes_root_version_1
// func ext_storage_changes_root_version_1(parent_hash int64) int64

// //go:wasm-module env
// //export ext_storage_clear_prefix_version_1
// func ext_storage_clear_prefix_version_1(prefix int64)

// //go:wasm-module env
// //export ext_storage_clear_version_1
// func ext_storage_clear_version_1(key_data int64)

// //go:wasm-module env
// //export ext_storage_get_version_1
// func ext_storage_get_version_1(key int64) int64

// //go:wasm-module env
// //export ext_storage_next_key_version_1
// func ext_storage_next_key_version_1(key int64) int64

// //go:wasm-module env
// //export ext_storage_read_version_1
// func ext_storage_read_version_1(key int64, value_out int64, offset int32) int64

// //go:wasm-module env
// //export ext_storage_root_version_1
// func ext_storage_root_version_1() int64

// //go:wasm-module env
// //export ext_storage_set_version_1
// func ext_storage_set_version_1(key int64, value int64)

// //go:wasm-module env
// //export ext_storage_exists_version_1
// func ext_storage_exists_version_1(key int64) int32

// /*
// 	Trie: Interface that provides trie related functionality
// */
// func ext_trie_blake2_256_ordered_root_version_1(data int64) int32
