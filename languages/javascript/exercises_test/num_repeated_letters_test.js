import { numRepeatedLetters } from '../exercises_itemized/num_repeated_letters';


test('numRepeatedLetters test', function(t) {
    t.plan(3);
    t.equal(numRepeatedLetters("aaaaabb"), 2);
    t.equal(numRepeatedLetters("abcdabcdeeffgghh"), 8);
    t.equal(numRepeatedLetters("abcdefghijklmnopqrstuvwxyz"), 0);
});