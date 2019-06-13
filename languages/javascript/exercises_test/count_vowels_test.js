import { countVowels } from '../exercises_itemized/count_vowels';


test('countVowels test', function(t) {
    t.plan(3);
    t.equal(countVowels("onomatopoeia"), {a: 2, e: 1, i: 1, o: 4});
    t.equal(countVowels("undulate"), {a: 1, e: 1, u: 2});
    t.equal(countVowels("Lisa left her job in the city"), {a: 1, e: 2, i: 3, o: 1});
});