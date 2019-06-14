import { fibonacci } from './exercises';


test('fibonacci test', function(t) {
    t.plan(5);
    t.equal(fibonacci(3), 2);
    t.equal(fibonacci(6), 8);
    t.equal(fibonacci(7), 13);
    t.equal(fibonacci(9), 34);
    t.equal(fibonacci(15, 610));
});