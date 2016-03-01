class Commodity(val abbr: String)

class Amount(
  val qty: BigDecimal,
  val cmdty: Commodity) {

  def this(qty: String, cmdty: String) {
    this(BigDecimal(qty), new Commodity(cmdty))
  }

  def +(that: Amount): Amount = {
    if(this.cmdty.abbr != that.cmdty.abbr) {
      throw new Exception("Illegal adding of incompatible Amounts")
    }
    new Amount(this.qty + that.qty, that.cmdty)
  }

  override def toString() = this.cmdty.abbr + " " + this.qty.toString
}

object App {
  def main(args: Array[String]): Unit = {
    var a1 = new Amount(BigDecimal("4.53"), new Commodity("$"))
    var a2 = new Amount("4.53", "$")
    var a3 = a1 + a2
    var a4 = a1 + a3 + new Amount(BigDecimal("4.00"), new Commodity("$"))
    println(a4)
  }
}
