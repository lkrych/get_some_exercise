import { timeConversion } from '../exercises_itemized/time_conversion';


test('timeConversion test', function(t) {
    t.plan(3);
    t.equal(timeConversion(15), "0:15");
    t.equal(timeConversion(150), "2:30");
    t.equal(timeConversion(360), "6:00");
});