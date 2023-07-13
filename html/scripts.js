var filtered_band_info

async function BandInfo(bandid) {
    /*let BandId = document.querySelector(".band_id"+bandid).id
    console.log(BandId)*/

    const url = '/bandinfo';
    /*const data = {
        username: 'johnDoe',
        email: 'john@example.com'
    };*/

    await fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/text'
        },
        body: JSON.stringify(bandid)
    })
        .then(response => response.json())
        .then(result => {
            filtered_band_info = result
            fetch_band_template(filtered_band_info)
            // Handle the response data
        })
        .catch(error => {
            console.error('Error:', error);
            // Handle errors
        });
}

function fetch_band_template(filtered_band_info) {
    const url_arist_page = '/artisttemplate.html';
    const targetElement = document.getElementsByTagName('body')[0];

    fetch(url_arist_page)
        .then(response => response.text())
        .then(html => {
            html = html.replace(searchValue = `.{{ImagePath}}`, replaceValue = filtered_band_info.Image)
            html = html.replace(searchValue = `.{{ArtistName}}`, replaceValue = filtered_band_info.Name)
            html = html.replace(searchValue = `.{{ArtistMembers}}`, replaceValue = filtered_band_info.Members)
            html = html.replace(searchValue = `.{{CreationDate}}`, replaceValue = filtered_band_info.CreationDate)
            html = html.replace(searchValue = `.{{FirstAlbumDate}}`, replaceValue = filtered_band_info.FirstAlbum)
            html = html.replace(searchValue = `.{{Locations}}`, replaceValue = filtered_band_info.Locations)
            targetElement.innerHTML = html;
        })
        .catch(error => {
            console.error('Error:', error);
            // Handle errors
        });
}