let tests = "test suite for last_two" >::: [
      "empty" >:: (fun _ -> assert_equal None (last_two []));
      "one"   >:: (fun _ -> assert_equal (None) (last_two [1]));
      "many"  >:: (fun _ -> assert_equal (Some ("c", "d")) (last_two [ "a"; "b"; "c"; "d"]));
]

let _ = run_test_tt_main tests