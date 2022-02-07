package main

func (p *plugin) clear() (undo bool, err error) {
	if undo = !p.Clear; undo {
		return
	}

	// 清理存储桶
	// client.Bucket.Get()

	return
}
