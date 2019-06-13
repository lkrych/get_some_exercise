import { reverse } from '../exercises_itemized/reverse';


test('reverse test', function(t) {
    t.plan(4);
    t.equal(reverse("abcdefg"), "gfedcba");
    t.equal(reverse("lalalala"), "alalalal");
    t.equal(reverse("racecar"), "racecar");
    t.equal(reverse("jabba the hut"), "tuh eht abbaj");
});