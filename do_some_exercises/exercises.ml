(* last returns the last element of a list *)

let rec last = function
  | [] -> None
  | x :: [] -> Some x
  | x :: tail -> last tail