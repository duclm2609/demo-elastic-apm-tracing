package vn.duclm.gateway.entity;

public class PaymentGatewayRes {
    private String fnc;
    private String pgUserCode;
    private String data;

    public String getFnc() {
        return fnc;
    }

    public void setFnc(String fnc) {
        this.fnc = fnc;
    }

    public String getPgUserCode() {
        return pgUserCode;
    }

    public void setPgUserCode(String pgUserCode) {
        this.pgUserCode = pgUserCode;
    }

    public String getData() {
        return data;
    }

    public void setData(String data) {
        this.data = data;
    }
}
