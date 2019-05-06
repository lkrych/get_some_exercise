(* kth_element returns the kth element of a list *)
let rec kth_element k = function
| [] -> None
| x::tail -> if k = 1 then Some x else kth_element (k - 1) tail
