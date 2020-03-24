package vn.duclm.demo.microinventoryservice.dto;

import lombok.Data;

@Data
public class InventoryReq {
    private String id;
    private String name;
    private String category;
    private String price;
}
