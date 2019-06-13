import { nearbyAZ } from '../exercises_itemized/nearby_az';


test('nearbyAZ test', function(t) {
    t.plan(5);
    t.equal(nearbyAZ("a"), false);
    t.equal(nearbyAZ("az"), true);
    t.equal(nearbyAZ("abz"), true);
    t.equal(nearbyAZ("abcdz"), false);
    t.equal(nearbyAZ("za"), false);
});