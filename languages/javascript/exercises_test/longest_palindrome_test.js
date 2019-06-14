import { longestPalindrome } from './exercises';


test('longestPalindrome test', function(t) {
    t.plan(3);
    t.equal(longestPalindrome("abcbd"), "bcb");
    t.equal(longestPalindrome("abba"), "abba");
    t.equal(longestPalindrome("abcbdeffe"), "effe");
});