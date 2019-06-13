import { caesarCipher } from '../exercises_itemized/caesar_cipher';


test('caesarCipher test', function(t) {
    t.plan(3);
    t.equal(caesarCipher(3, "abc"), "def");
    t.equal(caesarCipher(3, "abc xyz"), "def abc");
    t.equal(caesarCipher(29, "abc xyz"), "def abc");
});