@env-config
Feature: create a resources using collaborative posixfs 

  Background:
    Given user "Alice" has been created with default attributes


  Scenario: create folder
    Given the config "STORAGE_USERS_POSIX_WATCH_FS" has been set to "true"
    And user "Alice" has uploaded file with content "This is version 2" to "textfile.txt"
    When the administrator creates folder "myFolder" for user "Alice"
    Then the command should be successful
    And as "Alice" folder "/myFolder" should exist
    