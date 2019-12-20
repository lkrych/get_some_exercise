#include "gtest/gtest.h"
#include "exercises.h"
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
 


int main(int argc, char **argv) {
	testing::InitGoogleTest(&argc, argv);

	return RUN_ALL_TESTS();
}
