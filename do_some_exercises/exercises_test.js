var test = require('tape'); 
import { palindrome } from './exercises';


test('palindrome test', function(t) {
    t.plan(3);
    t.equal(palindrome("abc"), false);
    t.equal(palindrome("racecar"), true);
    t.equal(palindrome("z"), true);
});
 
