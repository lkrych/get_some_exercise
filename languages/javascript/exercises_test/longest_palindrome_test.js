import { longestPalindrome } from '../exercises_itemized/longest_palindrome';


test('longestPalindrome test', function(t) {
    t.plan(3);
    t.equal(longestPalindrome("abcbd"), "bcb");
    t.equal(longestPalindrome("abba"), "abba");
    t.equal(longestPalindrome("abcbdeffe"), "effe");
});