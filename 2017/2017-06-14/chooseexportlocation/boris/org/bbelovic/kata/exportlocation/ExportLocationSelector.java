package org.bbelovic.kata.exportlocation;

import java.util.List;

import static java.util.Comparator.comparing;


public class ExportLocationSelector {

        public ExportLocation chooseExportLocation(List<ExportLocation> exportLocations, final long shareId) {
            return exportLocations.stream()
                .filter(el -> el.getSharedId() == shareId)
                .filter(el -> !el.isAdmin())
                .sorted(comparing(ExportLocation::isPreferred).reversed())
                .findFirst().orElseThrow(RuntimeException::new);
    }
}
