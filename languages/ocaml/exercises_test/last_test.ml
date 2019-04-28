
let tests = "test suite for last" >::: [
      "empty" >:: (fun _ -> assert_equal None (last []));
      "one"   >:: (fun _ -> assert_equal (Some 1) (last [1]));
      "many"  >:: (fun _ -> assert_equal (Some "d")(last [ "a"; "b"; "c"; "d"]));
]

let _ = run_test_tt_main tests
