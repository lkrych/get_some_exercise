let tests = "test palindrome?" >::: [
  "empty list"   >:: (fun _ -> assert_equal true (palindrome []));
  "non palindrome" >:: (fun _ -> assert_equal false (palindrome ["h"; "e"; "l"; "l"; "o"]));
  "palindrome"     >:: (fun _ -> assert_equal true (palindrome ["r"; "a"; "c"; "e"; "c"; "a"; "r"]));
]

let _ = run_test_tt_main tests