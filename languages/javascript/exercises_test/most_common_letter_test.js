import { mostCommonLetter } from '../exercises_itemized/most_common_letter';


test('mostCommonLetter test', function(t) {
    t.plan(3);
    t.equal(mostCommonLetter("aaaaaaaaaaaabbbbbcccdssdfasfaasdfs"), "a");
    t.equal(mostCommonLetter("bbbbbbbbbbbbbbbbadfgadfgpioadfjgpadoifb"), "b");
    t.equal(mostCommonLetter("cccccccccpoiasdfgjad;foigjawccccwerac"), "c");
});