import { numRepeatedLetters } from './exercises';


test('numRepeatedLetters test', function(t) {
    t.plan(3);
    t.equal(numRepeatedLetters("aaaaabb"), 2);
    t.equal(numRepeatedLetters("abcdabcdeeffgghh"), 8);
    t.equal(numRepeatedLetters("abcdefghijklmnopqrstuvwxyz"), 0);
});