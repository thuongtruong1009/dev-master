interface Point {
    x: number;
    y: number;
}

const regularPoint: Point = {
    x: 10,
    y: 10,
};

type Readonly<T> = {
    readonly [P in keyof T]: T[P];
};

const readonlyPoint: Readonly<Point> = {
    x: 10,
    y: 10,
};

regularPoint.x = 20;
console.log(regularPoint);

// readonlyPoint.x = 20;
console.log(readonlyPoint);

function movedPoint(p: Point, dx: number, dy: number): Point {
    return {
        x: p.x + dx,
        y: p.y + dy,
    };
}

const movedRegularPoint = movedPoint(regularPoint, 3, 4);
console.log(movedRegularPoint);

// const movedReadonlyPoint = movedPoint(readonlyPoint, 3, 4);
