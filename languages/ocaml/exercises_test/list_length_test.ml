let tests = "test suite for length" >::: [
      "empty" >:: (fun _ -> assert_equal 0 (length []));
      "one"   >:: (fun _ -> assert_equal 1 (length [1]));
      "many"  >:: (fun _ -> assert_equal 4 (length [ "a"; "b"; "c"; "d"]));
]

let _ = run_test_tt_main tests