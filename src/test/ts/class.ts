export default class Car {
  private vendor: string;
  private model: string;
  private releaseDate: Date;

  constructor(vendor: string, model: string) {
    this.vendor = vendor;
    this.model = model;
    this.releaseDate = new Date();
  }

  public get = (): string => {
    return `${this.releaseDate} - ${this.vendor} ${this.model}`;
  };
}
