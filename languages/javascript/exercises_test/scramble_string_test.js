import { scrambleString } from './exercises';


test('scrambleString test', function(t) {
    t.plan(2);
    t.equal(scrambleString("abcd", [3, 1, 2, 0]), "dbca");
    t.equal(scrambleString("markov", [5, 3, 1, 4, 2, 0]), "vkaorm");
});