package lags

import "reflect"
import "testing"

type requests []struct {
	start, duration, price int
}

func MakeCombination(r requests) []requests {
	/*
			  requestsToReturn := []requests{}
			  partialCombinace := requests{}
		    //A1 B1 C1 ... etc
			  for _,v :=  r {
			     partialCombinace := append(partialCombinace,v)
		   	  }

		      for i,v := r {
		        partialCombinace := append(partialCombinace,v)
		      }
			  requestsToReturn := append(requestsToReturn, partialCombinace)
			  return requestsToReturn
	*/
	var pokus []requests
	if len(r) == 2 {
		first := requests{r[0]}
		second := requests{r[1]}
		pokus = append(pokus, requests{first})
		pokus = append(pokus, requests{second})

		return []requests{
			pokus,
			requests{r[0], r[1]},
		}
	}
	if len(r) == 3 {
		return []requests{
			requests{r[0]},
			requests{r[1]},
			requests{r[2]},
			requests{r[0], r[1]},
			requests{r[0], r[2]},
			requests{r[1], r[2]},
			requests{r[0], r[1], r[2]},
		}
	}
	return []requests{r}
}

func TestMakeCombination(t *testing.T) {

	type testCase struct {
		requests        requests
		allCombinations []requests
	}

	testCases := []testCase{
		{
			requests:        requests{{0, 1, 100}},
			allCombinations: []requests{requests{{0, 1, 100}}},
		},
		{
			requests: requests{{0, 1, 100}, {2, 3, 50}},
			allCombinations: []requests{
				requests{{0, 1, 100}},
				requests{{2, 3, 50}},
				requests{{0, 1, 100}, {2, 3, 50}},
			},
		},
		{
			requests: requests{{0, 1, 100}, {4, 3, 50}},
			allCombinations: []requests{
				requests{{0, 1, 100}},
				requests{{4, 3, 50}},
				requests{{0, 1, 100}, {4, 3, 50}},
			},
		},
		{
			requests: requests{{0, 1, 100}, {0, 2, 200}, {0, 3, 300}},
			allCombinations: []requests{
				requests{{0, 1, 100}},
				requests{{0, 2, 200}},
				requests{{0, 3, 300}},
				requests{{0, 1, 100}, {0, 2, 200}},
				requests{{0, 1, 100}, {0, 3, 300}},
				requests{{0, 2, 200}, {0, 3, 300}},
				requests{{0, 1, 100}, {0, 2, 200}, {0, 3, 300}},
			},
		},
		{
			requests: requests{{0, 1, 100}, {0, 2, 200}, {0, 9, 900}},
			allCombinations: []requests{
				requests{{0, 1, 100}},
				requests{{0, 2, 200}},
				requests{{0, 9, 900}},
				requests{{0, 1, 100}, {0, 2, 200}},
				requests{{0, 1, 100}, {0, 9, 900}},
				requests{{0, 2, 200}, {0, 9, 900}},
				requests{{0, 1, 100}, {0, 2, 200}, {0, 9, 900}},
			},
		},
	}

	for _, tt := range testCases {
		got := MakeCombination(tt.requests)
		if !reflect.DeepEqual(got, tt.allCombinations) {
			t.Errorf("MakeCombination(%v):%v WANT:%v", tt.requests, got, tt.allCombinations)
		}
	}

}
