import XCTest

import clientTests

var tests = [XCTestCaseEntry]()
tests += clientTests.allTests()
XCTMain(tests)
