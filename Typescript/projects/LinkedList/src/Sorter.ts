interface Sortable {
    length: number;
    compare(leftindex: number, rightindex: number): boolean;
    swap(leftindex: number, rightindex: number): void;
    sort(): void;
}

export default abstract class Sorter implements Sortable {
    abstract length: number;
    abstract compare(leftindex: number, rightindex: number): boolean;
    abstract swap(leftindex: number, rightindex: number): void;

    sort(): void {
        const { length } = this;
        for (let i = 0; i < length; i++) {
            for (let j = 0; j < length - i - 1; j++) {
                if (this.compare(j, j + 1)) {
                    this.swap(j, j + 1);
                }
            }
        }
    }
}
