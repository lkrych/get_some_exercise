import { powerOfTwo } from './exercises';


test('powerOfTwo test', function(t) {
    t.plan(6);
    t.equal(powerOfTwo(3), false);
    t.equal(powerOfTwo(6), false);
    t.equal(powerOfTwo(8), true);
    t.equal(powerOfTwo(512), true);
    t.equal(powerOfTwo(4194304), true);
    t.equal(powerOfTwo(4194305, false))
});