// Copyright 2017-2018 DERO Project. All rights reserved.
// Use of this source code in any form is governed by RESEARCH license.
// license can be found in the LICENSE file.
// GPG: 0F39 E425 8C65 3947 702A  8234 08B2 0360 A03A 9DE8
//
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND ANY
// EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL
// THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
// PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT,
// STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF
// THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package checkpoints

// generate checksums automatically from file mainnet_checksums.dat

// generate blank file if no testnet checkpoints
//go:generate sh -c "echo package checkpoints > testnet_checksums.go"
//go:generate sh -c "if [ ! -s testnet_checksums.dat ]; then echo var testnet_checksums_base64 = \"\" >> testnet_checksums.go; fi;"
//go:generate sh -c " if [  -s testnet_checksums.dat ]; then echo var testnet_checksums_base64 = \\\x60 >> testnet_checksums.go; fi"
//go:generate sh -c "if [  -s testnet_checksums.dat ]; then base64 -w 80 <  testnet_checksums.dat >> testnet_checksums.go;fi"
//go:generate sh -c "if [  -s testnet_checksums.dat ]; then echo \\\x60 >> testnet_checksums.go;fi "

//go:generate sh -c "echo // Code generated by go generate DO NOT EDIT. > mainnet_checksums.go"
//go:generate sh -c "echo // This file contains all the mainnet checksums > mainnet_checksums.go"
//go:generate sh -c "echo // please read checkpoints.go comments > mainnet_checksums.go"
//go:generate sh -c "echo package checkpoints >> mainnet_checksums.go"
//go:generate sh -c "echo  >> mainnet_checksums.go"
//go:generate sh -c "echo var mainnet_checksums_base64 = \\\x60 >> mainnet_checksums.go" 
//go:generate sh -c "base64 -w 80 <  mainnet_checksums.dat >> mainnet_checksums.go"
//go:generate sh -c "echo  \\\x60 >> mainnet_checksums.go" 

import "fmt"

//import "bytes"
import "io/ioutil"
import "path/filepath"
import "encoding/base64"

import "github.com/romana/rlog"
import log "github.com/sirupsen/logrus"
import "github.com/armon/go-radix"

//import "github.com/deroclassic/deroclassic_suite/crypto"
import "github.com/deroclassic/deroclassic_suite/globals"

// this file handles and maintains checkpoints which are used by blockchain to skip some checks on known parts of the chain
// the skipping of the checks can be disabled by command line arguments

//var mainnet_checkpoints_height uint64 = uint64(len(mainnet_checkpoints) / 32) // each checkpoint is 32 bytes in size
//var testnet_checkpoints_height uint64 = uint64(len(testnet_checkpoints) / 32) // each checkpoint is 32 bytes in size

var mainnet_checksums_height uint64 = uint64(len(mainnet_checksums) / 32) // each checksum is 32 bytes in size
var testnet_checksums_height uint64 = uint64(len(testnet_checksums) / 32) // each checksum is 32 bytes in size

var checksum_tree *radix.Tree

func init() {
	//	rlog.Tracef(1, "Loaded %d checkpoints for mainnet hardcoded", mainnet_checkpoints_height)
	//	rlog.Tracef(1, "Loaded %d checkpoints for testnet hardcoded", testnet_checkpoints_height)

	rlog.Tracef(1, "Loaded %d checksums for mainnet hardcoded", mainnet_checksums_height)
	rlog.Tracef(1, "Loaded %d checksums for testnet hardcoded", mainnet_checksums_height)
	checksum_tree = radix.New()
}

var mainnet_checksums = load_base64(mainnet_checksums_base64)
var testnet_checksums = load_base64(testnet_checksums_base64)

func load_base64(input string) []byte {
    
    data, err := base64.StdEncoding.DecodeString(input)
    if err != nil {
        rlog.Tracef(1, "Loaded checksums failed base64 decoding input length %d", len(input))
    }
    
    return data
    
}

// load checkpoints from the data directory
// a line should be printed on console when we are doing this
func LoadCheckPoints(logger *log.Entry) {
	/*
		if globals.IsMainnet() { // load mainnet checkpoints
			data_loaded, err := ioutil.ReadFile(filepath.Join(globals.GetDataDirectory(), "mainnet_checkpoints.dat"))
			if err == nil && uint64(len(data_loaded)%32) == 0 {
				mainnet_checkpoints = data_loaded
				mainnet_checkpoints_height = uint64(len(mainnet_checkpoints) / 32)
				rlog.Tracef(1, "Loaded %d checkpoints for mainnet from file  mainnet_checkpoints.dat", mainnet_checkpoints_height)
				logger.Infof("Loaded %d checkpoints for mainnet from file  mainnet_checkpoints.dat", mainnet_checkpoints_height)
			} else {
				rlog.Warnf("Loading checkpoints for mainnet from file mainnet_checkpoints.dat failed err %s len %d", err, uint64(len(mainnet_checkpoints)))
				if len(data_loaded) > 0 {
					logger.Warnf("Loading checkpoints for mainnet from file mainnet_checkpoints.dat failed err %s len %d", err, uint64(len(mainnet_checkpoints)))
				}
			}
		}

		if !globals.IsMainnet() { // load testnet checkpoints
			data_loaded, err := ioutil.ReadFile(filepath.Join(globals.GetDataDirectory(), "testnet_checkpoints.dat"))
			if err == nil && uint64(len(data_loaded)%32) == 0 {
				testnet_checkpoints = data_loaded
				testnet_checkpoints_height = uint64(len(testnet_checkpoints) / 32)
				rlog.Tracef(1, "Loaded %d checkpoints for testnet from file  testnet_checkpoints.dat", testnet_checkpoints_height)
				logger.Info(1, "Loaded %d checkpoints for testnet from file  testnet_checkpoints.dat", testnet_checkpoints_height)
			} else {
				rlog.Warnf("Loading checkpoints for testnet from file testnet_checkpoints.dat failed err %s len %d", err, uint64(len(testnet_checkpoints)))
				if len(data_loaded) > 0 {
					logger.Warnf("Loading checkpoints for testnet from file testnet_checkpoints.dat failed err %s len %d", err, uint64(len(testnet_checkpoints)))
				}
			}
		}
	*/
	if globals.IsMainnet() { // load mainnet checksum
		data_loaded, err := ioutil.ReadFile(filepath.Join(globals.GetDataDirectory(), "mainnet_checksums.dat"))
		if err == nil && uint64(len(data_loaded)%32) == 0 {
			mainnet_checksums = data_loaded
			mainnet_checksums_height = uint64(len(mainnet_checksums) / 32)
			rlog.Tracef(1, "Loaded %d checksums for mainnet from file  mainnet_checksum.dat", mainnet_checksums_height)
			logger.Infof("Loaded %d checksums for mainnet from file  mainnet_checksum.dat", mainnet_checksums_height)
		} else {
			rlog.Warnf("Loading checksums for mainnet from file mainnet_checksum.dat failed err %s len %d", err, uint64(len(mainnet_checksums)))
			if len(data_loaded) > 0 {
				logger.Warnf("Loading checksums for mainnet from file mainnet_checksum.dat failed err %s len %d", err, uint64(len(mainnet_checksums)))
			}
		}
	}

	if !globals.IsMainnet() { // load testnet checksum
		data_loaded, err := ioutil.ReadFile(filepath.Join(globals.GetDataDirectory(), "testnet_checksums.dat"))
		if err == nil && uint64(len(data_loaded)%32) == 0 {
			testnet_checksums = data_loaded
			testnet_checksums_height = uint64(len(testnet_checksums) / 32)
			rlog.Tracef(1, "Loaded %d checksums for testnet from file  testnet_checksum.dat", testnet_checksums_height)
			logger.Infof("Loaded %d checksums for testnet from file  testnet_checksum.dat", testnet_checksums_height)
		} else {
			rlog.Warnf("Loading checksums for testnet from file testnet_checksums.dat failed err %s len %d", err, uint64(len(testnet_checksums)))
			if len(data_loaded) > 0 {
				logger.Warnf("Loading checksums for testnet from file testnet_checksums.dat failed err %s len %d", err, uint64(len(testnet_checksums)))
			}
		}
	}

	// lets build a trie for checksums for faster lookup
	checksum_tree = radix.New()
	if globals.IsMainnet() {
		for i := uint64(0); i < mainnet_checksums_height; i++ {
			checksum_tree.Insert(string(mainnet_checksums[32*i:32*(i+1)]), nil)
		}
	} else {
		for i := uint64(0); i < testnet_checksums_height; i++ {
			checksum_tree.Insert(string(testnet_checksums[32*i:32*(i+1)]), nil)
		}
	}

	logger.Debugf("Succesfully built checksum tree count %d", checksum_tree.Len())

}

/*
// gives length of currently available checkpoints
func Length() uint64 {
	switch globals.Config.Name {
	case "mainnet":
		return mainnet_checkpoints_height
	case "testnet":
		return testnet_checkpoints_height
	default:
		return 0
	}

	// we can never reach here
	//return 0
}
*/
func ChecksumLength() uint64 {
	switch globals.Config.Name {
	case "mainnet":
		return mainnet_checksums_height
	case "testnet":
		return testnet_checksums_height // no checksums for testnet
	default:
		return 0
	}

	// we can never reach here
	//return 0
}

/*
// tell whether a checkpoint is known in the current selected network
func IsCheckPointKnown(hash crypto.Hash, height uint64) (result bool) {

	var known_hash crypto.Hash

	switch globals.Config.Name {
	case "mainnet":
		if height < mainnet_checkpoints_height {
			copy(known_hash[:], mainnet_checkpoints[32*height:])
			if known_hash == hash {
				result = true
				return
			}
		}

	case "testnet":
		if height < testnet_checkpoints_height {
			copy(known_hash[:], testnet_checkpoints[32*height:])
			if known_hash == hash {
				result = true
				return
			}
		}

	default:
		panic(fmt.Sprintf("Unknown Network \"%s\"", globals.Config.Name))
	}
	return
}

*/
// tell whether a checkpoint is known in the current selected network
func IsCheckSumKnown(hash []byte) (result bool) {
	switch globals.Config.Name {
	case "mainnet":
		if _, ok := checksum_tree.Get(string(hash)); ok {
			return true
		}

	case "testnet":
		if _, ok := checksum_tree.Get(string(hash)); ok {
			return true
		}
		return false

	default:
		panic(fmt.Sprintf("Unknown Network \"%s\"", globals.Config.Name))
	}
	return
}