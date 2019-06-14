import { nthPrime } from './exercises';


test('nthPrime test', function(t) {
    t.plan(5);
    t.equal(nthPrime(3), 2);
    t.equal(nthPrime(6), 8);
    t.equal(nthPrime(7), 13);
    t.equal(nthPrime(9), 34);
    t.equal(nthPrime(15, 610));
});