#include "string"

TEST(ReverseStrTest, Blahblah) {
	std::string x = "blahblah";
    std::string reversedx ="halbhalb";
	reverseStr(x);
	EXPECT_TRUE(x == reversedx);
}

TEST(ReverseStrTest, Racecar) {
	std::string x = "racecar";
    std::string reversedx ="racecar";
	reverseStr(x);
	EXPECT_TRUE(x == reversedx);
}