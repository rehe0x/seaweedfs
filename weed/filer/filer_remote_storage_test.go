package filer

import (
	"github.com/chrislusf/seaweedfs/weed/pb/filer_pb"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFilerRemoteStorage_FindRemoteStorageClient(t *testing.T) {
	conf := &filer_pb.RemoteConf{
		Name: "s7",
		Type: "s3",
	}
	rs := NewFilerRemoteStorage()
	rs.storageNameToConf[conf.Name] = conf

	rs.mapDirectoryToRemoteStorage("/a/b/c", "s7")

	_, found := rs.FindRemoteStorageClient("/a/b/c/d/e/f")
	assert.Equal(t, true, found, "find storage client")

	_, found2 := rs.FindRemoteStorageClient("/a/b")
	assert.Equal(t, false, found2, "should not find storage client")

	_, found3 := rs.FindRemoteStorageClient("/a/b/c")
	assert.Equal(t, false, found3, "should not find storage client")

	_, found4 := rs.FindRemoteStorageClient("/a/b/cc")
	assert.Equal(t, false, found4, "should not find storage client")
}