exports.LikesCount = async () => {
    const resp = await fetch('/api/likes')
    const data = await resp.json()
    return data.likes
}

exports.LikesIncrement = async () => {
    const resp = await fetch('/api/likes', {method: 'POST'})
    console.log(await resp.text())
}