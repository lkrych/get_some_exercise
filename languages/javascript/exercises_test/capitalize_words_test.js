import { capitalizeWords } from './exercises';


test('capitalizeWords test', function(t) {
    t.plan(3);
    t.equal(capitalizeWords("hi there"), "Hi There");
    t.equal(capitalizeWords("My name is nine"), "My Name Is Nine");
    t.equal(capitalizeWords("My name is nine not 9"), "My Name Is Nine Not 9");
});