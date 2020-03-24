package vn.duclm.demo.microinventoryservice.dto;

import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class InventoryRes {
    private String id;
    private String name;
    private String category;
    private String price;
}
