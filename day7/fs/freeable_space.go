package fs

const (
	maxDirSize          = 100001
	totalSpace          = 70000000
	updateRequiredSpace = 30000000
)

func TotalFreeableSpace(dir *Directory) int {
	total := 0
	for _, f := range dir.content {
		if d, ok := f.(*Directory); ok {
			if d.Size() < maxDirSize {
				total += d.Size()
			}
			total += TotalFreeableSpace(d)
		}
	}
	return total
}

func SmallestSubdirectoryToDelete(rootDir *Directory) int {
	totalSize := rootDir.Size()
	remainingSpace := totalSpace - totalSize
	requiredSpace := updateRequiredSpace - remainingSpace
	return findMinDirectorySizeAboveNumber(rootDir, requiredSpace)
}

func findMinDirectorySizeAboveNumber(dir *Directory, s int) int {
	min := 999999999
	for _, f := range dir.content {
		if d, ok := f.(*Directory); ok {
			if d.Size() >= s {
				if d.Size() < min {
					min = d.Size()
				}
			}
			subMin := findMinDirectorySizeAboveNumber(d, s)
			if subMin < min && subMin >= s {
				min = subMin
			}
		}
	}
	return min
}
