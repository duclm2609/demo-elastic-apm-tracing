package vn.duclm.demo.microinventoryservice;

import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import vn.duclm.demo.microinventoryservice.dto.InventoryRes;

@RestController
@RequestMapping("/")
@Slf4j
public class InventoryController {

    @Autowired
    private InventoryService inventoryService;

    @GetMapping("/{id}")
    public InventoryRes getInventoryDetail(@PathVariable("id") String id) {
        log.info("receive get inventory detail request");
        return this.inventoryService.getInventoryDetail(id);
    }
}
