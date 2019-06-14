import { longestWord } from './exercises';


test('longestWord test', function(t) {
    t.plan(3);
    t.equal(longestWord("one two three"), "three");
    t.equal(longestWord("The bronchioles or bronchioli are the passage ways by which air passes through the nose or mouth to the alveoli"), "bronchioles" );
    t.equal(longestWord("lalalalalala tada wa wa ma ma"), "lalalalalala");
});