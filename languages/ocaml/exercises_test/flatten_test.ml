let tests = "test flatten" >::: [
  "empty list" >:: (fun _ -> assert_equal [] (flatten []));
  "non-nested list" >:: (fun _ -> assert_equal [1; 2; 3;] (flatten [One 1; One 2; One 3;]));
  "singly-nested list" >:: (fun _ -> assert_equal [1; 2; 3;](flatten [One 1; Many [One 2;]; One 3;]));
  "doubly-nested list" >:: (fun _ -> assert_equal [1; 2; 3;] (flatten [One 1; Many [ Many [One 2;]; One 3;]]));
]

let _ = run_test_tt_main tests
