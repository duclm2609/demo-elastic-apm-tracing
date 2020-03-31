package vn.duclm.demo.microinventoryservice;

import lombok.extern.log4j.Log4j2;
import org.apache.logging.log4j.message.StringMapMessage;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.PathVariable;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import vn.duclm.demo.microinventoryservice.dto.InventoryRes;

@RestController
@RequestMapping("/")
@Log4j2
public class InventoryController {

    @Autowired
    private InventoryService inventoryService;

    @GetMapping("/{id}")
    public InventoryRes getInventoryDetail(@PathVariable("id") String id) {
        log.info(new StringMapMessage()
                .with("message", "This is hello message")
                .with("type", "HTTP_REQ"));
        return this.inventoryService.getInventoryDetail(id);
    }
}
