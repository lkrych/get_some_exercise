#include <string>

TEST(ReverseStrTest, blahblah) {
	std::string x = "blahblah";
    std::string reversedx ="halbhalb";
	EXPECT_EQ(reverseStr(x), reversedx);
}

TEST(ReverseStrTest, racecar) {
	std::string x = "racecar";
    std::string reversedx ="racecar";
	EXPECT_EQ(reverseStr(x), reversedx);
}