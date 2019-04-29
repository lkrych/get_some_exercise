(*
let rec kth_element k = function
  [] -> None
  | k' :: ks' -> if k = 1 then Some k' else kth_element (k-1) ks'  
*)