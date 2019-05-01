let tests = "test suite for reverse" >::: [
  "empty" >:: (fun _ -> assert_equal [] (reverse []));
  "one"   >:: (fun _ -> assert_equal [1] (reverse [1]));
  "many"  >:: (fun _ -> assert_equal [3; 2; 1;] (reverse [1; 2; 3;]));
]

let _  = run_test_tt_main tests