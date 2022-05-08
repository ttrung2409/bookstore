package com.bookstore.ecommerce.app.domain.data;

import lombok.Builder;
import lombok.Data;

@Builder
@Data
public class Book {
  private String id;
  private String title;
  private String subtitle;
  private String description;
  private String thumbnailUrl;
}
