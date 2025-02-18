/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.codegen.python;

import java.util.stream.IntStream;

/**
 * Helper class to print code.
 */
public class CodePrinter {

    private StringBuffer buffer = new StringBuffer();
    private int tabSize;

    private int intendationLvl = 0;

    public CodePrinter(int tabSize) {
        this.tabSize = tabSize;
    }

    public void startBlock() {
        this.intendationLvl += tabSize;
    }

    public void endBlock() {
        if (intendationLvl < tabSize) {
            throw new RuntimeException("Closing a Block which is not open!");
        }
        this.intendationLvl -= tabSize;
    }

    public void write(String s) {
        buffer.append(s);
    }

    public void startLine(String s) {
        writeIntendation();
    }

    public void endLine() {
        buffer.append("\n");
    }

    public void writeLine(String s) {
        writeIntendation();
        buffer.append(s);
        buffer.append("\n");
    }

    private void writeIntendation() {
        // Write the intendation
        IntStream.range(0, intendationLvl).forEach(i -> buffer.append(" "));
    }

    public String getCode() {
        return buffer.toString();
    }
}
