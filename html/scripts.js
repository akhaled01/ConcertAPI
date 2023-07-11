function BandInfo(bandid) {
    let BandId = document.querySelector(".band_id"+bandid).id
    console.log(BandId)
    const get = async (url, params) => {
        const response = await fetch(url, {
            method: 'GET',
            body: JSON.stringify(params),
            headers: {
                'Content-type': 'application/json; charset=UTF-8',
            }
        })
    
        const data = await response.json()
    
        return data
    }
    
    // Then use it like so with async/await:
    (async () => {
        const response = await fetch('/artists?' + new URLSearchParams({
            id: BandId
        }))
    
        const data = await response.json()
    
        console.log(data)
    })()
}