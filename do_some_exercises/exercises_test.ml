open OUnit2;;

open Exercises;;

let tests = "test suite for kth_element" >::: [
      "empty" >:: (fun _ -> assert_equal None (kth_element 1 []));
      "one"   >:: (fun _ -> assert_equal (None) (kth_element 2 [1]));
      "many"  >:: (fun _ -> assert_equal (Some ("c")) (kth_element 3 [ "a"; "b"; "c"; "d"]));
]

let _ = run_test_tt_main tests
 
