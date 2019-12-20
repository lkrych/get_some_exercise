#easily generate exercise, test and solution files

def get_file_suffix(lang):
  if lang == "python":
    return ".py"
  elif lang == "go":
    return ".go"
  elif lang == "c":
    return ".c"
  elif lang == "cpp":
    return ".cpp"
  elif lang == "elixir":
    return ".ex"
  elif lang == "ocaml":
    return ".ml"
  elif lang == "javascript":
    return ".js"
  else:
    raise Exception("This tool does not support that language")



lang = input("What language do you want to add an exercise, solution and test for? ")
exercise_name = input("What is the name of your exercise? ")

try:
  suffix = get_file_suffix(lang)
except Exception as error:
  print(error)
  exit(0)

exercise_path = "/{lang}/exercises_itemized/{ex}{suffix}".format(lang=lang, ex=exercise_name, suffix=suffix)
solution_path = "/{lang}/exercises_solutions/{ex}_soln{suffix}".format(lang=lang, ex=exercise_name, suffix=suffix)
test_path = "/{lang}/exercises_test/{ex}_test{suffix}".format(lang=lang, ex=exercise_name, suffix=suffix)

print("This will create an exercise named {ex}\n a solution named {sol}\n and a test named {test} in the {lang} folder".format(ex=exercise_path, sol=solution_path, test=test_path, lang=lang))

ok = input("Is this okay? y or n: ")
if ok == "y":
  #make files
  path_prefix = "./../languages"
  ex_file = open(path_prefix + exercise_path, "w")
  ex_file.close()
  soln_file = open(path_prefix + solution_path, "w")
  soln_file.close()
  test_file = open(path_prefix + test_path, "w")
  test_file.close()
else:
  print("Okay, I'm going to die now. Please just run this script again if you need to make some files")
  exit(0)
  