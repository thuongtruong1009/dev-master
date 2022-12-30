import { IMappable } from './interface';

export default class Map {
    private googleMap: google.maps.Map;
    private divId: string;

    constructor(divId: string) {
        this.divId = divId;
        this.googleMap = new google.maps.Map(
            document.querySelector(`#${divId}`)!,
            {
                zoom: 1,
                center: {
                    lat: 0,
                    lng: 0,
                },
            },
        );
    }

    addMarker(mappable: IMappable): void {
        const marker = new google.maps.Marker({
            map: this.googleMap,
            position: {
                lat: mappable.location.lat,
                lng: mappable.location.lng,
            },
        });

        marker.addListener('click', () => {
            const infoWindow = new google.maps.InfoWindow({
                content: mappable.markerContent(),
            });
            infoWindow.open(this.googleMap, marker);
        });
    }
}
