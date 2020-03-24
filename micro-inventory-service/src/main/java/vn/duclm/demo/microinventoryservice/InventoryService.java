package vn.duclm.demo.microinventoryservice;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.web.client.RestTemplateBuilder;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestTemplate;
import vn.duclm.demo.microinventoryservice.dto.InventoryRes;
import vn.duclm.demo.microinventoryservice.dto.PriceRes;

@Service
@Slf4j
public class InventoryService {

    private final RestTemplate restTemplate;

    public InventoryService(RestTemplateBuilder restTemplateBuilder) {
        this.restTemplate = restTemplateBuilder.build();
    }

    public InventoryRes getInventoryDetail(String id) {
        try {
            //Fake long running task
            Thread.sleep(700);
        } catch (InterruptedException e) {
            e.printStackTrace();
        }

        // Call price service
        log.info("getting product price from Price service");
        PriceRes price = restTemplate.getForObject("http://micro-price:8080/price/{id}", PriceRes.class, id);

        return InventoryRes.builder()
                .id(id)
                .name("Mercedes-Maybach S-Class Saloon")
                .category("Car")
                .price(price.getPrice())
                .build();
    }
}
