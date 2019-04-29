(* last returns the last element of a list *)
let rec last = function
| [] -> None
| a::[] -> Some a
| h::t -> last t
